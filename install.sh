#!/bin/bash

go install

GOPATH_BIN="$(go env GOPATH)/bin"

cat > ~/Library/LaunchAgents/com.user.privacyshare.plist << EOF
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple Computer//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
   <key>Label</key>
   <string>com.user.privacyshare</string>
   <key>ProgramArguments</key>
   <array>
      <string>$GOPATH_BIN/privacyshare</string>
   </array>
   <key>RunAtLoad</key>
   <true/>
</dict>
</plist>
EOF

launchctl bootstrap gui/$(id -u) ~/Library/LaunchAgents/com.user.privacyshare.plist

echo "PrivacyShare installed and started successfully!"
