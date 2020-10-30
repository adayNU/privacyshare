# PrivacyShare

The idea of this tool is to quietly remove tracking parameters from URLs you copy to the clipboard.
Especially in cases where you aren't explicitly copying text (i.e. clicking "copy to clipboard"), tools
tend to like to inject additional tracking parameters to understand who you're sharing things with.

#### Example:
When sharing a link in Spotify like this:  

![img](https://i.imgur.com/c9VWWRw.png)

The resulting string on your clipboard would be:
```
https://open.spotify.com/track/2UUVwDVZYR5StS7Si0SxrP?si=aa3kRCBySzqM3dPucu5GEg
```

However, if running PrivacyShare, it would instead simply be:
```
https://open.spotify.com/track/2UUVwDVZYR5StS7Si0SxrP
```

In its current form this is just a POC, however, I may try to make this more encompassing / usable by
non-technical users.
 