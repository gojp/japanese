[![Build Status](https://travis-ci.org/gojp/japanese.svg?branch=master)](https://travis-ci.org/gojp/japanese) [![go report card](http://goreportcard.com/badge/gojp/japanese)](http://goreportcard.com/gojp/japanese)
# Japanese

A Go (golang) package for generating the different possible word forms and conjugations in Japanese. 

This is still under heavy development. It's basically only an experiment at this stage. However, later extensions might also include basic grammar parsing, analysis and translation. 

## Where this is aiming

A `GetRoot` function that takes a string and returns one or more root forms, along with the rule used to transform it to root form. For example:

```
>>> roots := japanese.Analyze("食べない")
>>> root := roots[0]
>>> fmt.Println(root.Root, root.DictionaryForm, root.Rule)
食べ 食べる negative
```

## Examples
For the most comprehensive examples, see the words_test.go file. For the impatient, here's an example of usage right now; it's bound to change in the (very near) future:

    func main() {
        taberu := Verb{Verb{Word{"食べる", "たべる"}}
        kau := Verb{Verb{Word{"買う", "かう"}}
        matsu := Verb{Word{"待つ", "まつ"}}

        fmt.Println(taberu.Negative()) // prints 食べない
        fmt.Println(kau.Negative())    // prints 買わない
        fmt.Println(matsu.Negative())  // prints 待たない
    }
