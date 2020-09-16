# What is this?
This is a rewrite of the python example that waveshare provides for their 1in44 display hat

Currently only the buttons are functional after making sure this python runs on startup

```py


```

This project is held back by this error

https://github.com/stianeikeland/go-rpio/issues/66 

Because for some reason, even touching the SPI functions causes pins 21 and 20 (key1 and key2) to be toggled to `0`, which is their value when pressed, and it is unchangable unless you rerun the original python or reboot your py