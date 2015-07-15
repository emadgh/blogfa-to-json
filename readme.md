# Blogfa to JSON
I didn't work on this so much, the code might not looks good but it works *Good*,
Just wrote it to fetch all my brother blog posts on Blogfa.com to a new place

I used it to fetch more than 600 post from blogfa.

Might be usefull for others,
Use it on your own.

## What would it fetch?
well, it will fetch all post, tags, their comments


## Output scheme
It would be looks like this
```js
[
  {
    "Title": "",
    "Content": "",
    "Date":"2014-01-10T00:00:00+03:30",
    "Tags":["tag","tag",...],
    "Comments":[
      {
        "Name":"عماد قاسمی",
        "Comment":"سلام",
        "Date":"2014-01-11T07:21:00+03:30"
      },
      ...
    ]
  },
  ...
]
```

## Suggestion
It's better to edit your Theme and make it confortable for script to find elements

For example you can use this format
```html
<div class='post'>
  <div class='title'><span class='beg'></span><a href="<-PostLink->" class='posttitle'><-PostTitle-></a><span class='end'></span></div>
  <div class='postbody'>

    <div class='postcontent'><-PostContent-></div>

    <BlogPostTagsBlock>
      <div class='posttags'>برچسب‌ها:
      <BlogPostTags separator=", ">
        <a href="<-TagLink->" class='tagname'><-TagName-></a>
      </BlogPostTags></div>
    </BlogPostTagsBlock>

    <BlogExtendedPost>
      <a href="<-PostLink->">ادامه مطلب</a>
    </BlogExtendedPost>

    <div class="postdesc">
      <a href="<-PostLink->" title="لينك ثابت">+</a> نوشته شده توسط <-PostAuthor-> در <span class='postdate'><-PostDate-></span> و ساعت
      <-PostTime-> | <BlogComment><span dir="rtl"><script type="text/javascript">GetBC(<-PostId->);</script></span></BlogComment>
    </div>
  </div>
</div>
```
