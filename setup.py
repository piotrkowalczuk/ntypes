# -*- coding: utf-8 -*-
"""
:copyright: (c) 2016 by Piotr Kowalczuk
:license: MIT, see LICENSE for more details.
"""
from setuptools import setup

with open('VERSION.txt', 'r') as content_file:
    version = content_file.read()


    def readme():
        """print long description"""
        with open('README.md') as f:
            return f.read()


    setup(
        name='protobuf-ntypes',
        version=version[1:],
        description='protobuf nilable data structures',
        url='https://github.com/piotrkowalczuk/ntypes',
        author='Piotr Kowalczuk',
        author_email='p.kowalczuk.priv@gmail.com',
        license='MIT',
        packages=['ntypes'],
        install_requires=[
            'protobuf'
        ],
        # Classifiers help users find your project by categorizing it.
        #
        # For a list of valid classifiers, see https://pypi.org/classifiers/
        classifiers=[  # Optional
            'Development Status :: 5 - Production/Stable',
            'Intended Audience :: Developers',
            'Topic :: Software Development :: Libraries :: Application Frameworks',
            'License :: OSI Approved :: MIT License',

            'Programming Language :: Python',
            'Programming Language :: Python :: 2',
            'Programming Language :: Python :: 2.7',
            'Programming Language :: Python :: 3',
            'Programming Language :: Python :: 3.4',
            'Programming Language :: Python :: 3.5',
            'Programming Language :: Python :: 3.6',
        ],
        zip_safe=False,
        keywords=['protobuf', 'data-structures', 'grpc'],
        download_url='https://github.com/piotrkowalczuk/ntypes/archive/%s.tar.gz' % version.rstrip(),
    )
