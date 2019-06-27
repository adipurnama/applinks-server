# Serving Universal Links / App Links from Go

Simple go webserver for iOS Universal Links & Android App Links asset.

## Prerequisite

Do the following :

* Install Go version >= 1.12.x
* Read the documentation for :
  * [Android App Links](https://developers.google.com/digital-asset-links/v1/getting-started)
  * [iOS Universal Links](https://developer.apple.com/ios/universal-links/)

## How to

1. Change the content of well-known directory
   * apple-app-site-association for iOS
   * assetlinks for Android
2. ```go run main.go```

 by default, the server runs on port 4000
