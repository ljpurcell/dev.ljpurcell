---
title: "SICP-Go Notes: Introduction"
slug: sicp-go-intro
tags: go, sicp
---


TLDR:
1. SICP is a computer science text book held in the highest regard
3. I am going to be working through SICP using Go and posting my notes
4. The goal is to extract the insights of SICP, enhance my understanding of Go, and take another step on the path of becoming a master programmer

Every field has its canonical texts. In computer science, _Structure and Interpretation of Computer Programs_ is one of these books. It may be _the_ book. Commonly known as SICP, or the “Wizard Book,” it is notoriously difficult as a supposed introductory text.

What gives SICP such a reputation is that the notorious difficulty is accompanied by [glowing reviews from industry legends](https://www.amazon.com/review/R403HR4VL71K8/). It is claimed that no book compares when it comes to generating insights into programming and computer science.

The excellent blogger, Gwern, opened a series of posts on the book like this:

{.quote}
> What can be said about it that hasn’t been said a thousand times before? It is perhaps the single greatest computer science textbook … It is renowned for its imaginative exercises and mind-blowing techniques. So it’s pretty good.
>
> {.citation}
> [SICP Introduction, Gwern](https://gwern.net/sicp/introduction)




The book was originally published in 1984, with a second edition being released in 1996. In the subsequent two and a half decades, the relevance of the book has been questioned _ad nauseum_. Uncountable articles, forum discussions, and questions posed to communities and experts alike have broached this topic. As with any classic text, the answer is not blatantly clear.

Comparably, the ancient Greeks had much to say about ethics, and their contributions have been formative to current conceptions, but they said all of it before we had modern science – which entails heliocentrism, evolution, the germ theory of disease, among many other [viewquake-level](https://www.lesswrong.com/posts/zCf3pnQmMhyEK8Lit/on-viewquakes#:~:text=a%20viewquake%20is%20an%20%22insight%20that%20dramatically%20changes%20one%27s%20worldview%2C%20making%20one%20see%20the%20world%20in%20a%20new%20way.%22) discoveries. 

In the same way, Abelson and the Sussman – the authors of SICP – were writing from a much earlier point along the computer science timeline. Since then, new concepts have been coined; we know more about programming language design; hardware and resource constraints have evolved; and the process of constructing programs, large or small, has changed. 

Things are different. 

## Is SICP still relevant?

Yes.

In an act of outright defiance to questions of this kind, a [third, adapted version of SICP](https://sourceacademy.org/sicpjs/index) was released in 2022. Nearly a full four-decades after the original; a relative eternity given the field itself is not that much older. This time, using a more modern programming language. The original versions of the book used a dialect of Lisp, a functional language, called Scheme (renamed to Racket). The adapted version has made use of a much more recognisable language to contemporary eyes: JavaScript.

This changes both nothing and everything.

In relation to the former, a programming language is just a tool. Some, certainly, are better suited to specific problem domains. But here – where the language is simply being used as a medium to teach fundamental programming concepts – both are considered more than sufficient. Furthermore, JavaScript shares many similarities in design with Scheme/Racket. 

In effect, changing the language is more superficial than fundamental.

But, also, there is profound importance to this as well. This adapted version – with the superficial modern icing thrown atop the same SICP cake is a bold statement of, _“Yes! You better believe the content is still relevant.”_

## My experience with SICP

SICP is one of those book that has stared at me from my bookshelf, and desk, for years. I have a copy of both the JS and Scheme versions. However, owning a book is a far, far cry from having read it. And having read a book is far, far cry from having understood it.

I have attempted to do this, at least twice, previously. 

Sometime around the end of 2020, I attempted to work my way through the (Scheme version of the) book. I made it to page 25. Of which, a good portion included introductory remarks and an outline of the book. Not exactly the crucial computer science content the book is renowned for and the reason I was reading it. I also skipped a significant number of the exercises or attempted them yet could not solve them. I desperately wanted to acquire the wisdom contained within the covers – but I was not ready.

Then, in mid-2022, the book was calling me again. This time, I thought I might give the modernised JavaScript version a go. While Scheme was chosen in the original books due to its ease of learning and minimal primitive complexity – allowing for maximal expression on behalf of the programmer – I still felt that it left me a half-step behind. Instead, I reasoned, I may be stacking the odds ever-so-slightly more in my favour if I could read the material and attempt the same problems when it was framed in terms of JavaScript. I made it some 90 pages. Better, but not far enough.


## This series

It doesn't take a genius to know what I am going to say at this point. But, I will say it anyway.

I am going back to SICP. I am going to read it. And, more importantly, I am going to understand it. I will achieve this understanding by doing the exercises and writing summaries (which I will post here); following the model of _"see one, do one, teach one."_

My language of choice, however, is going to be Go (Golang). 

Go will (obviously) not translate perfectly to either the Scheme or JS versions of the book -- and that is fine. My goal in working through the book is to extract the general programming wisdom of the text, and use it as an exercise to increase my understanding of Go. Where any major divergences occur between Go and (Scheme or JS), I will do my best to note and illustrate them

This is important as JS was specifically chosen for the adapted version of the book due to its commonalities with Scheme/Racket, such as lexical scoping, higher-order functions and tail-call optimisation. 

We will cover what each of these is, and more, throughout this series. And it will help us understand Go at a deeper level by comparing and contrasting it against the languages SICP uses.
