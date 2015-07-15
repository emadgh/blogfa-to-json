# Blogfa to json

I didn't work on this so much,
Just wrote it fast to fetch all my brother blog posts on Blogfa.com to a new place

I used it to fetch more than 600 post from blogfa.

Might be usefull for others,
Use it on your own.


## suggestion
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
