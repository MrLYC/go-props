# coding: utf-8

import argparse
import os
import unittest
import shutil
from collections import defaultdict
from subprocess import check_output, Popen, PIPE


class TestParser(unittest.TestCase):
    target = os.getenv("TARGET", "./bin/go-props")
    testdata = os.getenv("TESTDATA", "testdata")

    def gen_code(self, file):
        return check_output([self.target, "-f", file], shell=False).strip()

    def read_code(self, file):
        with open(file, "rt") as fp:
            return fp.read().strip()

    def check_build(self, *files):
        bindir, _ = os.path.split(self.target)
        p = Popen(["go", "build", "-o", os.path.join(bindir, "simba.go")], stdin=PIPE)
        codes = os.linesep.join([self.read_code(f) for f in files])
        p.communicate(codes)
        self.assertEqual(p.poll(), 0)

    def test_data(self):
        cases = defaultdict(lambda: dict.fromkeys([".in", ".out"]))
        for f in os.listdir(self.testdata):
            name, ext = os.path.splitext(f)
            case = cases[name]
            if ext not in case:
                continue
            case[ext] = os.path.join(self.testdata, f)

        for case in cases.values():
            in_ = case[".in"]
            out = case[".out"]
            if not (in_ and out):
                continue

            self.assertEqual(self.gen_code(in_), self.read_code(out), in_)
            self.check_build(in_, out)


if __name__ == '__main__':
    unittest.main()
