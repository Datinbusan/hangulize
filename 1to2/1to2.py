# -*- coding: utf-8 -*-
from __future__ import print_function, unicode_literals

import io
import sys

import babel
import hangulize


def quote(x):
    return '"%s"' % x


class Section(object):

    def __init__(self, name):
        self.name = name
        self.pairs = []

    def put(self, k, *vs):
        self.pairs.append((k, map(unicode, vs)))

    def draw(self, sep, quote_keys=False):
        pairs = self.pairs[:]

        if not pairs:
            return ''

        if quote_keys:
            pairs = [(quote(k), vs) for k, vs in pairs]

        key_width = max(len(k) for k, vs in pairs)
        tmpl = '{0:%ds} {1} {2}' % key_width

        buf = io.StringIO()

        buf.write(self.name)
        buf.write(':\n')

        for k, vs in pairs:
            buf.write('    ')
            buf.write(tmpl.format(k, sep, quote(vs[0])))
            for v in vs[1:]:
                buf.write(', ')
                buf.write(quote(v))
            buf.write('\n')

        buf.write('\n')
        return buf.getvalue()


def main(argv):
    try:
        code = argv[1]
    except IndexError:
        print('Usage 1to2.py LANG')
        raise SystemExit(2)

    lang = hangulize.get_lang(code)
    locale = babel.Locale(lang.iso639_1)

    vars_ = []
    for attr in dir(lang.__class__):
        if attr.startswith('_'):
            continue
        if hasattr(lang.__class__.__bases__[0], attr):
            continue
        vars_.append(attr)
    if lang.vowels:
        vars_.append('vowels')

    sec = Section('lang')
    sec.put('id', code)
    sec.put('codes', lang.iso639_1, lang.iso639_3)
    sec.put('english', locale.get_language_name('en_US'))
    sec.put('korean', locale.get_language_name('ko_KR'))
    sec.put('script', '???')
    print(sec.draw('='), end='')

    sec = Section('config')
    sec.put('author', '???')
    sec.put('stage', '???')
    sec.put('markers', *lang.__tmp__)
    print(sec.draw('='), end='')

    sec = Section('macros')
    if lang.vowels:
        sec.put('@', '<vowels>')
    print(sec.draw('=', quote_keys=True), end='')

    sec = Section('vars')
    for var in vars_:
        sec.put(var, *getattr(lang, var))
    print(sec.draw('=', quote_keys=True), end='')


if __name__ == '__main__':
    main(sys.argv)
