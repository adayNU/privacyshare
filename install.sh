go install

echo "<?xml version=\"1.0\" encoding=\"UTF-8\"?>
<!DOCTYPE plist PUBLIC \"-//Apple Computer//DTD PLIST 1.0//EN\" \"http://www.apple.com/DTDs/PropertyList-1.0.dtd\">
<plist version=\"1.0\">
<dict>
   <key>Label</key>
   <string>com.user.privacyshare</string>
   <key>ProgramArguments</key>
   <array>
      <string>$GOPATH/bin/privacyshare</string>
   </array>
   <key>RunAtLoad</key>
   <true/>
</dict>
</plist>
" > ~/Library/LaunchAgents/com.user.privacyshare.plist

launchctl load ~/Library/LaunchAgents/com.user.privacyshare.plist

