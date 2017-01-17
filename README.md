# Simple Notifier

This is just a simple telegram notifier to notify if a proccess has been completed. Actually, you can convert this 
code using your favorite or fluent programming language you have. 

Sometime I need to leave compiling process while the proccess could be spent to many time. for Example, 
I need to build an Image docker remotely. It will take time too long for complex Dockerfile. So I just need
to create a bash file which there's some code inside to build image and then I run something like this.

$ /tmp/builder.sh && sleep 5 && notify "building is done" 



## Compile
$ cd /path/of/code
$ go install app

## HOW TO USE
$ bin/app "whatever the message you want to push to your telegram"




twiter: @tiduronline
