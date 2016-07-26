if [ -d ./bin ]; then
    echo "Cleaning bin directory..."
    rm -rf ./bin/*
fi

for GOOS in darwin windows linux; do
  for GOARCH in 386 amd64; do
    echo "Building $GOOS/$GOARCH..."
    FILENAME="pokevision"
    if [ "windows" == $GOOS ]; then
        FILENAME="pokevision.exe"
    fi
    GOOS=$GOOS GOARCH=$GOARCH go build -o ./bin/$GOOS/$GOARCH/$FILENAME ./
    if [ "windows" == $GOOS ]; then
        zip -rjX ./bin/$GOOS-$GOARCH.zip ./bin/$GOOS/$GOARCH/
    else
        tar -C ./bin/$GOOS/$GOARCH/ -cvzf ./bin/$GOOS-$GOARCH.tar.gz .
    fi
  done
done

echo "Building linux/arm..."
GOOS=linux GOARCH=arm go build -o ./bin/linux/arm/pokevision ./
tar -C ./bin/linux/arm/ -cvzf ./bin/linux-arm.tar.gz .
