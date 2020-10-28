#!/usr/bin/env bash

# print usage info
print_usage () {
  echo -e "\e[33mUsage:\e[0m"
  echo -e "  $0 [options]"
  echo -e ""
  echo -e "\e[33mOptions:\e[0m"
  echo -e "  \e[32m-h\e[0m    Print this help message"
  echo -e "  \e[32m-t\e[0m    Tag for versionning (ex: \"1.0.0\")"
  echo -e "  \e[32m-p\e[0m    Platform to build (ex: \"linux/amd64\"), can be used several times."
}

# print error message
print_error () {
    echo -e ""

    echo -e "\e[41m  $(printf '%*s' ${#1} "")  \e[0m"
    echo -e "\e[97;41m  $1  \e[0m"
    echo -e "\e[41m  $(printf '%*s' ${#1} "")  \e[0m"

    echo -e ""
}

# define build version (timestamp by default)
TAG=$(date -u +%Y%m%d%H%M%S)

# define available platforms
AVAILABLE_PLATFORMS=("darwin/amd64" "linux/amd64")

# parse options and arguments
while getopts ": h t: p:" OPTION; do
    case $OPTION in
        h)
            print_usage
            exit
            ;;

        t)
            TAG=("$OPTARG")
            ;;

        p)
            # check if platform is available
            if [[ ! " ${AVAILABLE_PLATFORMS[@]} " =~ " ${OPTARG} " ]]; then
                print_error "Platform \"$OPTARG\" is not available"
                exit 1
            fi
            PLATFORMS+=("$OPTARG")
            ;;

        \?)
            print_error "Option \"-$OPTARG\" is not defined"
            exit 1
            ;;
    esac
done

# use available platforms when no one is provide
if [ ${#PLATFORMS[@]} == 0 ]; then
    PLATFORMS=("${AVAILABLE_PLATFORMS[@]}")
fi

# define path of current script directory
DIR="$(cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd)"

for PLATFORM in "${PLATFORMS[@]}"
do
    # get OS and arch information from platform
    PLATFORM_SPLIT=(${PLATFORM//\// })
    OS=${PLATFORM_SPLIT[0]}
    ARCH=${PLATFORM_SPLIT[1]}

    # init output name
    NAME='aergie-'$OS'-'$ARCH

    # build binaries
    echo -n "Building $NAME...  "

    env GOOS=$OS GOARCH=$ARCH go build \
        -ldflags "-X 'main.version=$TAG'" \
        -o "$DIR/bin/$OS/$NAME" \
        $DIR/cmd/aergie/main.go

    if [ $? -ne 0 ]; then
        echo "An error has occurred! Aborting the script execution..."
        exit 1
    fi

    # update execution permission
    chmod +x $DIR/bin/$OS/$NAME

    echo "Done!"
done
