# Get HTML Attribute

Given a HTML node, get value of given attribute. If HTML is not be able to parse, return empty string. If attribute does not exist, return empty string.

## For example

### getAttr(<tag attr1="v1" attr2="v2"></tag>, attr2)

```
"v2"
```

### getAttr(<tag attr1="v1" attr2="v2">some string</tag>, attr2)

```
"v2"
```

### getAttr(<tag attr1="v1" attr2="v2" />, attr2)

```
"v2"
```

### getAttr(<tag attr1="v1" attr2="v2">, attr2)

```
""
```

### getAttr(<tag attr1="v1" attr2="v2></tag>, attr2)

```
""
```

### getAttr(<tag attr1="v1" attr2="v2></tag" />, attr2)

```
"v2></tag"
```

### getAttr(<tag attr1="v1" attr3="v2"></tag>, attr2)

```
""
```

### getAttr(<tag attr1="v1" attr2="v2" attr2="v3"></tag>, attr2)

```
"v2"
```

### getAttr(<tag attr1="v1" Attr2="V2"></tag>, attr2)

```
"V2"
```

## Join Code Challenge

[Code Challenge on Facebook](https://www.facebook.com/groups/202249393918924/)