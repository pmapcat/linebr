## `linebr` will insert a new line after provided text

### Rationale

There are a lot of files that don't have line breaks, 
they make Emacs/Repl go bad for some reason I don't know. 

Grep and Ag do not work nicely with file without linebreaks. 


This is a dead simple and fast tool to insert line breaks
after a certain text pattern. 

Here are some examples how it may be used

### Install 

`go install github.com/pmapcat/linebr@latest`


### Example 

```
echo `echo '{:name "Автомобільні диски", :parent_id 4653880, :category_id 4624949} {:name "Плитка", :parent_id 4639424, :category_id 4625710} {:name "Посуд", :parent_id 2394287, :category_id 4626594}' | ./linebr -p "}" 

# will output 

{:name "Автомобільні диски", :parent_id 4653880, :category_id 4624949}
 {:name "Плитка", :parent_id 4639424, :category_id 4625710}
 {:name "Посуд", :parent_id 2394287, :category_id 4626594}


```

#### 150 mb one line length file will take this amount of time to finish


```
time cat data/test_huge.txt | ./linebr -p "|"  | tail
4a53b4e7-e85a-4d60-b50c-768a1243dd55|
753b6ca9-4c69-4907-869e-48c1ed900a56|
bed998f0-0adf-4655-97b5-c81db40b846b|
1495d592-4131-4fad-94d9-37416762e777|
d660d940-0099-4cc4-915a-bb6a5084dfa0|
7f67e3ac-eadd-4034-8053-9f78631ed24c|
5bd738cc-de71-45cc-b570-660ac353f882|
6c6eab8f-6413-4857-ba95-1461621f5813|
d694b546-40d3-4d56-8b5a-90d2810bfa3c|
7e266fc9-88dc-445f-a76d-d4666abab2c9|
cat data/test_huge.txt  0,00s user 0,18s system 22% cpu 0,785 total
./linebr -p "|"  0,67s user 0,22s system 113% cpu 0,785 total
tail  0,20s user 0,15s system 45% cpu 0,784 total
```


### TODO

* Make this faster
