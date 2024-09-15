---
title: How to Find a Software Project (for Learning)
slug: learning
tags: learning
---

Learning is hard. That statement stands on its own. I don't need to qualify it further by saying *"Learning the guitar is hard"* or *"Learning physics is hard"*. Learning entails new, and new is hard.

This post is about learning software engineering and development. There are **general** lessons here, but the lens we are taking is software engineering. 

## Just do projects

It is well-established (and very correct) lore in The Land of Programmers that Project Based Learning (PBL) is the way to go. It is the only reliable algorithm to get you from **here**, to **there**. One iteration is entirely insufficient, but... you can be sure, with high a high degree of confidence, that the distance between where you are and where you want to go will shrink as a result.

The same cannot be said for other learning methods.

I wont sit here and be evangelical about PBL, though. You've heard it before. You will hear it again. It's a bit like being told to exercise, save your money, and get enough sleep at night. It may be correct advice -- in that, if executed, it yields positive return -- but you **already know** to do it. So it's not good advice. Being told to do it doesn't help you.

## Ok, so what then?

All I want to do is tell you a bit about my first **real** attempt at PBL.

Possibly like you, I had the goal of wanting to get better as a programmer. I was writing a fair bit of code, mostly at uni and then some other exercises. Sporadic attempt at Leetcode; the first 2-3 sections of some book, tutorial or course. You know the stuff.

I knew I should be taking on a project. If not for the learning benefits but the sheer fact I had no portfolio. I mean, I had uni assignments I had completed. There was substantial amounts of code I had written that I could point at and say, _"Look, I did this"_ -- but there is a kind of pre-chewedness to uni assignments. I knew it and so would a prospective employer.

(Not saying these are worthless. Definitely have a time and a place. For example, I found [auditing CS61B](https://sp24.datastructur.es/policies#auditing-cs61b), UC Berkeley's data-structures and algorithms course, **very** beneficial. More on that another time, maybe.)

One of my rationalisations for why I hadn't taken on any projects was that I didn't know what to work on. The follow up advice to _"build projects"_ is _"build something you want, would use, or fixes a problem in your life."_ But I didn't have any problems! My coffee machine reliably produces coffee each morning. The weather app works on my phone. Anything I built would be a lesser version of something that already existed. It would be contrived **and** bad.

And here's the breakthrough that I finally made... are you ready? 

{.note}
> It's ok to make bad and contrived things. 

If you believe in [Sturgeon's Law](https://en.wikipedia.org/wiki/Sturgeon%27s_law), then that shouldn't really matter. To put it bluntly, 9 out of 10 things you make will **mostly** suck... but you'll **nail** 1 in 10. 

Or something like that. Sturgeon's Law is only sort-of a thing anyway. Don't take it too seriously.

The main point is give yourself permission to make something bad. Because it won't ultimately be bad. It will just **start** that way. Software is either **good** or **unfinished**. ([NeoVim](https://neovim.io/) is both, btw.) 

It's a trap to consider your new projects (started or even just contemplating), which skew far more towards unfinished, and then compare those to all the amazing software you see out in the world. The latter err much more towards good end of the spectrum (and therefore much further away from the unfinished end). They are far later in their development and, in all likelihood, have had **tons** of crappy code laid down previously that has since been deleted or refactored.

{.note}
> I will add, though, something you make may not be as bad as you think. I have come to learn that even poor but tailored software very often ends up being better than very polished yet overly generalised software.

Earlier I said that not having a project to work on was a rationalisation. It was. In hindsight, there were many opportunities staring me in the face. It's a bit like Enlightenment in that sense though. You can only the see the profundity of the mundane once you're ready.

## Giving contrived and bad a chance

One evening I was at a mate's place for our end of season social basketball team's breakup. We were having dinner and doing a vote count, as one of the girlfriend's had graciously set up a Google Form for us to use to cast votes after each game.

As we were talking it became known that actually collating all the responses had been difficult. Google Forms, in all its infinite glory, kind of sucked for how **we wanted to use it**. But it was the best alternative... or was it?

The penny finally dropped for me. *"I could make something for us, for next season"* I thought. "*That's just a TODO list in disguise, right?*". But I knew the Sacred Text forbade TODO list projects on a portfolio. I wouldn't be that silly.

So that's what I did.

I built a web app for our social basketball team and learned a TON from it. And if you are a novice programmer, wanting to get better at software engineering and development, or looking to get your first job, you can find a comparable project for yourself.

You may not feel ready, but I would encourage you to dive in. I kept convincing myself that I needed to learn more before I could build something, without a tutorial, from start to finish. But the truth was taking on something new was the only way to actually learn more.

I am not understating my knowledge at the time when I say it was limited to something like:
1. A bunch of Python syntax, but very minimal understanding of modules
2. OOP means representing things as objects (srsly? what else could they be? I have never seen anything that wasn't an object)
3. JavaScript apparently does all the stuff people care about -- but I don't get how websites are dynamic?
4. Apps are a collection of instructions computers run
5. Databases hold data
6. There's something about this HTTP (and REST?) that allows users to send messages about creating, updating and deleting... Ohhhh, I get it now... That's what CRUD stands for.

That was it. I mean it.

## How get better at programming

You don't learn programming by writing 10 line functions, in isolation, that are the solution to some sandboxed problem. 

You learn programming by writing your fifth function for the day, getting stuck for 45 minutes, only then catching the bug you wrote in **the very first one** -- which all the others depend on. Then to decide the 2nd and 4th functions probably should never have existed in the first place anyway because they are making debugging a NIGHTMARE. Becoming frustrated at all the wasted effort. Storming away, having some food, thinking about something else... and then having the required knowledge begin to coalesce, sparking something -- ever so faintly. You think a new thought you'd never had before. You run back to your keyboard. Now you understand how your representation of the problem was wrong or what the stack trace is telling you. 

(Even though it was printed right there  in front of your eyes **every** time you ran and re-ran it while you were stuck, hoping it would be fixed, but now you can **see** it.)

That's how you learn programming.

And this is normally touted as the benefit of learning by debugging. Overcoming stuckness. The price of struggle for the payment of knowledge. But I think that misses the mark just a bit. Getting stuck, and overcoming that is definitely important. But I would get stuck in my assignments for school, and think I was experiencing what was required. I wasn't, and maybe you're not either. 


## Proper stuck

The problem with trying to get better at programming from assignments or coding challenges, is that when someone has given you a problem on a platter, you know there **is** a solution even if you don't know **the** solution. 

You can definitely get stuck and not know how to proceed. But by the very nature of the problem -- being smaller, more self-contained -- you're far more likely to try and find the shortcut, algorithm, or specific syntax you need to make it all fall into place. You are at Position A and Position B is your destination. You just need to know **how** to get there.

What I'm talking about instead, though, is a stuck that's qualitatively different. Proper stuck. And, if you're a novice programmer reading this... if you're at the stage where you are reading blogs about programming, you're probably ready to be proper stuck.

Proper stuck is when you think you're at Position A, and you want to get to Position Z, and you think Position B is where you need to go next... but you're not even really sure about that... because nothing seems to work the way you expected. You take, what you think is, one step north and realise you've moved three steps east.

This is the indication that your model of the situation is wrong. And the only way to solve it is improve how you think about the problem and your program, as well as maybe the computer and software at large.

It's painful, but worthwhile. You end up reading the docs more intently; reading to understand, rather than simply proceed. You look at your stack traces as clues to unravel, rather than annoying and indecipherable output signalling something has gone wrong.

I found proper stuck to be far more useful for learning **what I actually wanted to learn**. It forced me to think more deeply about the problem and develop a comfort with the tools at my disposal.

## The takeaway

Become a better programmer by building a custom solution a unique problem. The problem doesn't need to be **overwhelmingly**
 unique that other software couldn't be used as a solution. You just need a problem that you think you can tailor a solution towards. Your solution can be contrived and bad. The key is to make it yours. Making decisions and taking wrong turns. Because when you do, you won't be able to Google the exact way to unstuck yourself. No one will have been in that exact situation before. Instead, you'll be forced to get your bearings and learn to navigate. 
 
 And you'll be much better for it.

