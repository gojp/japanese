Japanese
=========

A Go (golang) package for generating the different possible word forms and conjugations in Japanese. 

This is still under heavy development. It's basically only an experiment at this stage. However, later extensions might also include basic grammar parsing, analysis and translation. 

Where this is aiming
--------

A `GetRoot` function that takes a string and returns one or more root forms, along with the rule used to transform it to root form. For example:

```
>>> root, rule := japanese.GetRoot("食べない")
>>> fmt.Println(root, rule)
食べる "negative"
```

Examples
--------
For the most comprehensive examples, see the words_test.go file. For the impatient, here's an example of usage right now; it's bound to change in the (very near) future:

    func main() {
        ruverb := RuVerb{Verb{Word{"食べる", "たべる", "eat"}}}
        uverb := UVerb{Verb{Word{"買う", "かう", "buy"}}}
        uverb2 := UVerb{Verb{Word{"待つ", "まつ", "wait"}}}

        fmt.Println(ruverb.Negative()) // prints 食べない
        fmt.Println(uverb.Negative()) // prints 買わない
        fmt.Println(uverb2.Negative()) // prints 待たない
    }
