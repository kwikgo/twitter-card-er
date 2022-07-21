# twitter-card-er

Dynamically produce twitter cards, for fun

# Setup 

## Build compile 

Like any ... ach. Clone this repository, i.e. download the source.
Then compile it, easiest by typing `make`.
This will give you an executable binary.
Using `make install` you can install that binary into go's standard binary location (`~/go/bin`), which is used in below #Raspian #RaspberyPi example (i.e. see below: `/home/pi/go/bin/twitter-card-er`).

## Installation

Mine is running below nginx like this:

```
ssl_certificate /etc/letsencrypt/live/tube.hacker.ch/fullchain.pem; 
ssl_certificate_key /etc/letsencrypt/live/tube.hacker.ch/privkey.pem;
ssl_dhparam /etc/letsencrypt/ssl-dhparams.pem; # managed by Certbot

server {
	listen 443 ssl;
	root /var/www/html/hacker.ch;
	index index.html;
        server_name hacker.ch; # managed by Certbot
	charset UTF-8;
	location /schnoddelbotz/ {
	    autoindex on;
	}
	location / {
		# First attempt to serve request as file, then
		# as directory, then fall back to displaying a 404.
		try_files $uri $uri/ =404;
	}

	location ~* \.(?:jpg|jpeg|gif|png|ico|cur|gz|svg|svgz|mp4|ogg|ogv|webm|htc)$ {
		expires 1d;
		#access_log off;
		add_header Cache-Control "public";
	}

	location /twitter-card-er/ {
        # As not consumed, save it ...
		# proxy_set_header X-Forwarded-For $remote_addr;
		# proxy_set_header Host            $http_host;
		proxy_pass http://localhost:9911/;
	}
}

server {
	listen 443 ssl;
    # stuff for tube.hacker.ch - #VHOST / #SSL #SAN #Cert ...
```

This can be kept alive by systemd, like on my Raspi:

```
pi@raspberrypi:~ $ cat /etc/systemd/system/twitter-card-er.service 
[Unit]
Description=twitter-card-er on 9911
ConditionPathExists=/home/pi/go/bin/twitter-card-er
After=network.target
[Service]
Type=simple
User=pi
Group=pi
Environment="TCE_BASEURL=https://hacker.ch/twitter-card-er"
Environment="TCE_SITE=@JanHacker9"
Environment="TCE_TITLE=Hacker's feucht-fröhlicher Twitter-Kater. Card-Er, Sie, Sorry!"
WorkingDirectory=/tmp
ExecStart=/home/pi/go/bin/twitter-card-er
Restart=on-failure
RestartSec=10
StandardOutput=syslog
StandardError=syslog
SyslogIdentifier=appgoservice
[Install]
WantedBy=multi-user.target
pi@raspberrypi:~ $
```

There are currently no plans to further ease installation ;-)

# PEACE 

Fight for #FreeSpeech by supporting Julian Assange.
For example, by following [Julian Assange's wife on Twitter](https://twitter.com/StellaMoris1).

# FreeAssange #FreeAssange 

<iframe width="560" height="315" src="https://www.youtube.com/embed/RphXR6U7Ir4" title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe>
