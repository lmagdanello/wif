# WIF (Where is the file?)

I find it a bit boring that sometimes I need to do a ```pwd``` and copy the filename or try to remember the entire path I'm working on. 

I know readlink does this, but you need to pass more arguments in addition to your file/directory name.  

That's why I created the WIF: just tell it the file/directory you want to analyze, it will show you its absolute path and its type (regular file? directory? maybe a symbolic link?)

Like that:

![wif](./screenshot/wif.png)