You can get notify about signals the os sends to your
process. for example, Ctrl-C or Ctrl-Z sends SIGINT signal.
Here how to get notified about SIGINT signal.
This program creates a channel, uses signal,
Notify to tell Go runtime to send a message to the channel
when the OS sends a signal to the process the ^C above is the
keyboard command CTRL+C which sends the SIGINT signal interrupt