neww -n dev/go
splitw -h -p 20 -t 0
selectw -t last-window
selectp -t 0
send-keys -t 1 "export PS1=' $ '" Enter
send-keys -t 0 'cd $GOPATH/src' Enter
