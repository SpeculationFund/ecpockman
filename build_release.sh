version=`date -u +%Y%m%d`
path=".Build"
product="ecpockman"


function build() {
    # X86
    X86=(amd64 386)

    for arch in ${X86[@]}; do
        env GOARCH=${arch} go build -o ${path}/${product}_${arch}-${version} ${code}
    done

    # ARM
    env GOARCH=arm GOARM=7 go build -o ${path}/${product}_arm-${version} ${code}
}


function compress() {
    zip -r ${product}-${version}.zip ${path}
    rm -rf ${path}
}


build
compress






