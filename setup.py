from setuptools import setup
from subprocess import check_output

with open('VERSION.txt', 'r') as content_file:
    version = content_file.read()

    setup(
        name='protobuf-ntypes',
        version=version[1:],
        description='protobuf nilable data structures',
        url='http://github.com/piotrkowalczuk/ntypes',
        author='Piotr Kowalczuk',
        author_email='p.kowalczuk.priv@gmail.com',
        license='MIT',
        packages=['ntypes'],
        install_requires=[
            'protobuf'
        ],
        zip_safe=False,
        keywords=['protobuf', 'data-structures'],
        download_url='https://github.com/piotrkowalczuk/ntypes/archive/%s.tar.gz' % version.rstrip(),
      )