import setuptools

with open('README.md', 'r') as f:
    description = f.read()

setuptools.setup(
    name='protopb', 
    version='v0.0.1', 
    author='solost23', 
    author_email='tianyuanyuans@163.com', 
    description=description, 
    url='https://github.com/solost23/protopb/tree/master/gen/py3', 
    packages=setuptools.find_package(),
)