# ste
ste (simple template engine) is a simple wrapper to go template engine, you can use it like this :

    ste templateFileName outputFileName string1 string2

where the template ([using go template syntax](https://golang.org/pkg/text/template/)) receive string1 as value1, string2 as value2 (and so on...), and values contains a slice with string1 and string2
