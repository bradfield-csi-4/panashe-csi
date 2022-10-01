#!/usr/bin/env python3
import re

def sub(m):
    return f"0x{int(m.group(0)):0>2x}"

def find_and_replace_decimal(s):
    return re.sub(r"\d+", sub, s)

def rewrite():
    with open("totranslate") as f, open("totranslate_hexed", "w") as out:
        for l in f:
            out.write(find_and_replace_decimal(l))

if __name__ == "__main__":
    rewrite()
