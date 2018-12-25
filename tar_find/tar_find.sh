# 
# It will:
# 
# deal with files with spaces, newlines, leading dashes, and other funniness
# handle an unlimited number of files
# won't repeatedly overwrite your backup.tar.gz like using tar -c with xargs will do when you have a large number of files
# Also see:
# 
# GNU tar manual http://www.gnu.org/software/tar/manual/tar.html#SEC107
# How can I build a tar from stdin?, searc
# https://stackoverflow.com/questions/2597875/how-can-i-build-a-tar-from-stdin
# 
find . -type f -print0 | tar -czvf backup.tar.gz --null -T -



# solution #1
#Then tar with the -T option which allows it to take a list of file locations (the one you just created with find!)

find . -name "*.whatever" > yourListOfFiles
tar -cvf yourfile.tar -T yourListOfFiles

# solution #2
find . -type f | xargs -d "\n" tar -czvf backup.tar.gz

# solution #3
#//If you have multiple files or directories and you want to zip them into independent *.gz file you can do this. Optional -type f -atime
find -name "httpd-log*.txt" -type f -mtime +1 -exec tar -vzcf {}.gz {} \;