---
title: "SICP Notes: Introduction"
slug: sicp-intro
tags: js, ocaml, sicp
---




TLDR:
1. SICP is a computer science text book held in the highest regard
2. It focuses on program organisation, abstracting detail, and functional programming, among many other fundamental concepts in computer science
3. I've tried, and failed, to work through it multiple times
4. I am giving it another go and will be posting my notes, aiming to solve most exercise twice &hyphen; once using JavaScript and then again using OCaml, a more functional language

Every field has its canonical texts. In computer science, _Structure and Interpretation of Computer Programs_ is one of these books. It may be _the_ book. Commonly known as SICP, or the “Wizard Book,” it is notoriously difficult as a supposed introductory text.

What gives SICP such a reputation is that the notorious difficulty is accompanied by [glowing reviews from industry legends](https://www.amazon.com/review/R403HR4VL71K8/). It is claimed that no book compares when it comes to generating insights into programming and computer science.

The excellent blogger, Gwern, opened a series of posts on the book like this:

{.quote}
> What can be said about it that hasn’t been said a thousand times before? It is perhaps the single greatest computer science textbook … It is renowned for its imaginative exercises and mind-blowing techniques. So it’s pretty good.
>
> {.citation}
> [Gwern, SICP Introduction](https://gwern.net/sicp/introduction)

The book was originally published in 1984, with a second edition being released in 1996. In the subsequent two and a half decades, the relevance of the book has been questioned _ad nauseum_. Uncountable articles, forum discussions, and questions posed to communities and experts alike have broached this topic. As with any classic text, the answer is not blatantly clear.

Comparably, the ancient Greeks had much to say about ethics, and their contributions have been formative to current conceptions, but they said all of it before we had modern science – which entails heliocentrism, evolution, the germ theory of disease, and many more [viewquake-level](https://www.lesswrong.com/posts/zCf3pnQmMhyEK8Lit/on-viewquakes#:~:text=a%20viewquake%20is%20an%20%22insight%20that%20dramatically%20changes%20one%27s%20worldview%2C%20making%20one%20see%20the%20world%20in%20a%20new%20way.%22) discoveries. 

In the same way, Abelson and the Sussman – the authors of SICP – were writing at a (relatively) much earlier time. Since then, new concepts have emerged; we know more about programming language design; hardware and resource constraints have evolved; and the process of constructing programs, large or small, has changed. 

Things are different. 

Is the book still relevant today?

## Is SICP still relevant?

Yes.

In an act of outright defiance to questions of this kind, a [third, adapted version of SICP](https://sourceacademy.org/sicpjs/index) was released in 2022. Nearly a full four-decades after the original (an eternity in the field of software engineering and computer science). This time, using a more modern programming language. The original versions of the book used a dialect of Lisp - a functional language - called Scheme. The adapted version has made use of a much more recognisable language to contemporary eyes: JavaScript.

This changes both nothing and everything.

In relation to the former, a programming language is just a tool. Some, certainly, are better suited to specific problem domains. But here – where the language is simply being used as a medium to teach fundamental programming concepts – both are considered more than sufficient. Furthermore, JavaScript shares many similarities in design with Scheme. 

In effect, changing the language is more superficial than fundamental.

But, also, there is profound importance to this as well. This adapted version – with the superficial modern icing thrown atop the same SICP cake is a bold statement of, _“Yes! You better believe the content is still relevant.”_


## But why SICP?

Ok, so people like SICP, and it was deemed relevant enough to re-publish very recently. But what makes it good? What makes it a worthy investment?

I will answer those questions in an almost entirely deferential manner, using quotes from the foreword and preface of the book. This should give an idea of what the 500+ page book represents and aims to bestow upon pupils.

{.quote}
> ... the subject matter of this book involves us with three foci of phenomena: the human mind, collections of computer programs, and the computer. Every computer program is a model, hatched in the mind, of a real or mental process.
>
> {.citation}
> [Alan J. Perlis, Foreword - SICP (1984)](https://sourceacademy.org/sicpjs/foreword84#p2)

{.quote}
> Since large programs grow from small ones, it is crucial that we develop an arsenal of standard program structures of whose correctness we have become sure—we call them idioms—and learn to combine them into larger structures using organizational techniques of proven value. These techniques are treated at length in this book, and understanding them is essential to participation in the Promethean enterprise called programming. More than anything else, the uncovering and mastery of powerful organizational techniques accelerates our ability to create large, significant programs.
>
> {.citation}
> [Alan J. Perlis, Foreword - SICP (1984)](https://sourceacademy.org/sicpjs/foreword84#p3)

{.quote}
> A programmer should acquire good algorithms and idioms.
>
> {.citation}
> [Alan J. Perlis, Foreword - SICP (1984)](https://sourceacademy.org/sicpjs/foreword84#p6)

{.quote}
> First, we want to establish the idea that a computer language is not just a way of getting a computer to perform operations but rather that it is a novel formal medium for expressing ideas about methodology. Thus, programs must be written for people to read, and only incidentally for machines to execute.
>
> {.citation}
> [Harold Abelson and Gerald Jay Sussman, Preface - SICP (1984)](https://sourceacademy.org/sicpjs/prefaces96#p7)


{.quote}
> Underlying our approach to this subject is our conviction that "computer science" is not a science and that its significance has little to do with computers. The computer revolution is a revolution in the way we think and in the way we express what we think. The essence of this change is the emergence of what might best be called procedural epistemology—the study of the structure of knowledge from an imperative point of view, as opposed to the more declarative point of view taken by classical mathematical subjects. Mathematics provides a framework for dealing precisely with notions of "what is." Computation provides a framework for dealing precisely with notions of "how to."
>
> {.citation}
> [Harold Abelson and Gerald Jay Sussman, Preface - SICP (1984)](https://sourceacademy.org/sicpjs/prefaces96#p10)


{.quote}
> SICP was never about a programming language; it presents powerful, general ideas for program organization that ought to be useful in any language.
>
> {.citation}
> [Guy L. Steele Jr., Foreword - SICP:JS (2021)](https://sourceacademy.org/sicpjs/foreword02#p13)

Here we see numerous references to the idea of program organisation, the importance of programs being descriptive and comprehensible to people, how computation is about processes, and the language-agnostic approach the book takes to exploration of all these topics.


## My experience with SICP

SICP is one of those book that has stared at me from my bookshelf, and desk, for years. I have a copy of both the JavaScript and Scheme version. However, owning a book is a far, far cry from having read it. And having read a book is far, far cry from having understood it.

I have attempted to do this, at least twice, previously. 

Sometime around the end of 2020, I attempted to work my way through the (Scheme version) book. I made it to page 25. Of which, a good portion included introductory remarks and an outline of the book. Not exactly the crucial computer science content the book is renowned for and the reason I was reading it. I also skipped a significant number of the exercises or attempted them yet could not solve them. I desperately wanted to acquire the wisdom contained within the covers – but I was not ready.

Then, in mid-2022, the book was calling me again. This time, I thought I might give the modernised JavaScript version a go. While Scheme was chosen in the original books due to its ease of learning and minimal primitive complexity – allowing for maximal expression on behalf of the programmer – I still felt that it left me a half-step behind. Instead, I reasoned, I may be stacking the odds ever-so-slightly more in my favour if I could read the material and attempt the same problems when it was framed in terms of JavaScript. I made it some 90 pages. Better, but not far enough.


## This series

It doesn't take a genius to know what I am going to say at this point. But, I will say it anyway.

I am going back to SICP. I am going to read it. And, more importantly, I am going to understand it. I will achieve this understanding by doing the exercises and writing summaries (which I will post here); following the model of _"see one, do one, teach one."_

Ok, so here's my plan...

My goal, in a sense, is to work through both editions of the book. While the content is extremely comparable, the Scheme version **did need to be adapted** to be compatible with JavaScript (which is not as much of a pure functional language). As such, in order to try and understand the book at its core, and deeply internalise the lessons, I aim to solve the exercises in each book. Most of this will be doubling up, and that's ok. Being forced to re-express the idea/solution in a second time, in a different language, can only benefit me.

I will start with the JavaScript edition. Then, once I solve all the exercise within a section of the JavaScript version, I will re-implement the more crucial ones using OCaml (as opposed to Scheme). I've wanted to learn OCaml (and a highly functional language, in general) for some time. This seems like a good way to way to that, while gaining additional exposure to SICP concepts.

Alright, that's enough of that... Time to crack in!nd it will help us understand Go at a deeper level by comparing and contrasting it against the languages SICP uses.
