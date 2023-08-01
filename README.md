# go-medium2markdown
go-medium2markdown (re: medium to markdown) is simple tool to convert your medium post to markdown format

# How it works?
With Medium post url, we can add query param `format=json` to get json format of the post. Then, we parse the json to markdown text.

*Example*

|Original Post|https://purnasatria.medium.com/medium-content-type-cce8b8bdd0bb|
|JSON format|https://medium.com/@purnasatria/medium-content-type-cce8b8bdd0bb?format=json|

## Medium JSON format
There are 2 main object in JSON result, `paragraphs` and `markups` 

### Paragraph Type
|Type|Description|
|---|---|
|1|Basic|
|3|Big T/ Heading 1|
|4|Image|
|6|Quote|
|8|Code Block|
|9|Unorder List|
|10|Order List|
|11|Embed|
|13|Small T/ Heading 2|
|14|Embed Link|

### Markup Type
|Type|Description|
|---|---|
|1|Bold|
|2|Italic|
|3|Link or Mention|
|10| Highlight/ In-line code|

# How to install

# How to use
```bash
md2 <medium_url>
```

Alternatively you can add yaml config file to set markup symbol (right now just for section and italic)
```bash
md2 <medium_url> -c <config_file_path>
```


# What can't be converted right now
- Github Gist embed

# Roadmap
[] output to file 
[] add Hugo Frontmatter
[] download asset/images

