#!/usr/bin/env python3
import RPi.GPIO as GPIO
import time
import sys

def control_power(seconds=1):
    GPIO.setmode(GPIO.BCM)
    GPIO.setup(18, GPIO.OUT)
    
    try:
        GPIO.output(18, GPIO.HIGH)
        time.sleep(seconds)
        GPIO.output(18, GPIO.LOW)
    finally:
        GPIO.cleanup()

def power_on():
    """开机 - 短按电源键"""
    control_power(1)

def power_off():
    """关机 - 长按电源键（通常 4-10 秒）"""
    control_power(5)

def restart():
    """重启 - 先关机再开机"""
    control_power(5)  # 长按关机
    time.sleep(2)
    control_power(1)  # 短按开机

if __name__ == "__main__":
    if len(sys.argv) < 2:
        print("Usage: python3 power_control.py [on|off|restart]")
        sys.exit(1)
    
    action = sys.argv[1].lower()
    if action == "on":
        power_on()
    elif action == "off":
        power_off()
    elif action == "restart":
        restart()
    else:
        print("Invalid action. Use: on, off, or restart")
