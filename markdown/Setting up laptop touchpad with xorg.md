[//]: # (17-10-2021)
# Setting up laptop touchpad with xorg


---

If you want to setup your laptop's touchpad through xorg here is the process:

Install libinput with your package manager:

```
sudo pacman -S libinput
```

Add the following to /etc/X11/xorg.conf.d/30-touchpad.conf

```
Section "InputClass"
    Identifier "touchpad"
   	    Driver "libinput"
    MatchIsTouchpad "on"
    Option "tapping" "on"
    Option "AccelProfile" "adaptive"
    Option "TappingButtonMap" "lrm"
EndSection
```
