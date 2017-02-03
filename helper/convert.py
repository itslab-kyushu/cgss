#!/usr/bin/env python
"""Convert binary files contain prime numbers in a readable format.
"""
from __future__ import print_function
import glob
import itertools
import struct


def is_prime(x):
    """Return True if the given x is a prime number.

    It is a quick check the conversion works well.
    """
    if x == 2:
        return True
    if x < 2 or x & 1 == 0:
        return False
    return pow(2, x - 1, x) == 1


def main():
    """The main function.
    """
    for name in glob.iglob("258-*"):

        with open(name, "rb") as fp:

            res = 0
            for c in itertools.count():
                b = fp.read(1)
                if not b:
                    break
                n = struct.unpack("B", b)[0]
                res += n * 256**c
            if is_prime(res):
                print(res)
            else:
                raise RuntimeError("Non prime number is found.")


if __name__ == "__main__":
    main()
