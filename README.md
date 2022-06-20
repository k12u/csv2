# csv2

Make csv greppable

This is heavily inspired by [gron](https://github.com/tomnomnom/gron)

## Usage

For a csv file you want to search:
```
$ cat > test.csv
first_name,last_name,username
"Rob","Pike",rob
Ken,Thompson,ken
"Robert","Griesemer","gri"
^D
```

csv2 expands the contents into a flat structure and makes it easy to grep. 

```
$ ./csv2  < ~/test.csv 
/0/first_name = Rob
/0/last_name = Pike
/0/username = rob
/1/first_name = Ken
/1/last_name = Thompson
/1/username = ken
/2/first_name = Robert
/2/last_name = Griesemer
/2/username = gri
```

You can search by specific content instead of sequential numbering by giving the "index" column name (`username`):

```
$ ./csv2 -i username < test.csv
/rob/first_name = Rob
/rob/last_name = Pike
/rob/username = rob
/ken/first_name = Ken
/ken/last_name = Thompson
/ken/username = ken
/gri/first_name = Robert
/gri/last_name = Griesemer
/gri/username = gri
```

You can search by meaningful value specified by index:

```
$ ./csv2 -i username < test.csv | grep /ken | grep last_name
/ken/last_name = Thompson 
```
