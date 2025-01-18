# RPI Go Bot

## About

Prototype of controlling Black Gladiator - caterpillar chassis with Iduino L298N motor driver, Raspberry Pi Zero W and L7805 stabilizer for powering Pi Zero. Whole project is powered by 6 AA batteries. Controller By Dualshock 3. And I know, that Pi Zero is obviusly overkill for this project. Software is written in Golang.

## Requirements

* caterpillar chassis with two dc motors
* some prebuild L298N motor driver
* Raspberry Pi Zero W with Raspberry Pi OS Lite (project is using vsersion from November 19th 2024)
    * best will be version with soldered headers if you don't want to solder
* Electronic parts
    * L7805
    * two ceramic capatictors around 100 mF
    * radiator for L7805
    * 6 AA battery holder or 2x3 AA battery holder, or ...
    * some cables, protopyle cabels and goldpins
    * power switch
    * small breadboard

## Schematic

![schematic](/assets/circuit.png)

L8705 can be build on small breadboard, remember to add radiator.

## Rpi system setup

To be able to connect to Dualshock 3 you need to change bluetooth configuration.
In file `/etc/bluetooth/input.conf` uncomment and change `ClassicBondedOnly` to `false`. After change restart bluetooth `sudo systemctl restart bluetooth`

Now you are able to connect controller. Execute commane `bluetoothctl`. Connect Dualshock 3 via USB, immediately you will be asked to authorize service, then add device to be trusted. 
Now you can unconnect controller and click PS button, it should connect to rpi via bluetooth:
```
$ bluetoothctl 
Agent registered
[CHG] Controller B8:27:EB:6D:60:8D Pairable: yes
Authorize service
[agent] Authorize service 00001124-0000-1000-8000-00805f9b34fb (yes/no): yes
[CHG] Device 00:06:F7:4A:65:D0 UUIDs: 00001124-0000-1000-8000-00805f9b34fb
[bluetooth]# devices
Device 00:06:F7:4A:65:D0 Sony PLAYSTATION(R)3 Controller
[bluetooth]# trust 00:06:F7:4A:65:D0
[CHG] Device 00:06:F7:4A:65:D0 Trusted: yes
Changing 00:06:F7:4A:65:D0 trust succeeded
[CHG] Device 00:06:F7:4A:65:D0 Connected: yes
```

After this controller should connect automatically after pressing PS button.
