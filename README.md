# Project in Cloud Technologies
![Contributors](https://img.shields.io/badge/Contributors-4-green.svg) ![Build Status](https://img.shields.io/badge/Build-Running-green.svg) ![Build Status](https://img.shields.io/badge/Coverage-75.7-blue.svg) 


## Contents
* [Introduction](#introduction)
* [Description of project](#description-of-project) | [Original plan](#original-plan) - [Final Result](#final-result)
* [**How to use the bot**](#how-to-use-the-bot)
* [**Testing**](#testing)
* [Reflection](#reflection) | [What went well](#what-went-well) - [What went wrong](#what-went-wrong)
* [What we have learned](#what-we-have-learned)
* [Total work hours](#total-work-hours)
* [References](#references)

***

<!-- PROJECT LOGO -->
<br />
<p align="center">
  <a href="https://git.gvk.idi.ntnu.no/andmag/project_cloud_2019">
    <img src="images/discordbotgif.gif" alt="Logo" width="260" height="210">
  </a>
</p>

## Introduction
**Groupname:** `PBM` | **Project name:** `Operation Assist Me`

For this project we wanted make use of what we had learned previous in the course Cloud Computing to explore technology related to making discord bots. It seemed like a fun idea to actually see if we could use what we have learned so far in our study program to make something we would find useful and actually could use ourselves. 

**Members:** 
* `Herman Andersen Dyrkorn`
* `Anders Langlie`
* `Eirik Martin Danielsen`
* `Andrea Magnussen`

## Description of project
A dicord bot that provides functionality such as: 
* Giving a real-time weather forecast
* A bill tracker system for the user
* A reminder that notifies the user after a certain time
* A simple newsfeed using webhooks

### Original plan
Our original plan was to make a discord bot that provided some sort of useful functionality, without actually setting specific goals to what we would implement. However we knew we wanted to create a scalable application which would be relatively easy to downscope or add further functionality to. A lot of the time spent on planning the project was used on exploring the different API's and possibilities that were availible to us. We did eventually decide on a theme for the project - usefulness in everyday life. We established early that user instructions and user feedback were extremely important, and we wanted to make sure that the bot is easy to both use and understand. 

### Final result
The final result is a bot that provides different functionality such as a real-time weather forecast, a bill tracker system, a reminder function, help commands for every functionality and more. The bot ended up satisfying our requirement for it to be scalable, it is therefore quite simple to add new functionality. We also added a webhook to the channel just because we wanted to explore the possibilities for using webhooks in discord. The webhook is invoked by the use of the PaaS IFTTT. This platform allows the user to connect different online services together. In our case we connected an RSS feed from a subreddit to the Discord webhook. This was not a part of the original plan, but we scaled the project a bit when we got a clearer picture of what we were capable of getting done before the deadline of the project. We reached a test coverage percent of 75.7% and by using golangci-lint we have no errors.  

## `How to use the bot`

1. **You need to have a discord account**
2. **You need to click the link that invites you to the Discord Channel. This link is provided to you in the project submission.**

The channel consists of: 
* **General:** Here you can test the bot by writing commands. **_!command_** is the syntax for commands. **_!help_** is the command to get you started.
* **Webhook:** This is the webhook channel. Messages will be posted to this channel when a webhook is invoked.

Follow the instructions provided by the **_!help_** command and have fun!

## `Testing`

Since we have some personal keys in our code, we have decided not to provide it due to security reasons. We will instead provide a picture of our test coverage. The tests are also availible in the tests folder. If you really want to actually run the tests, feel free to contact us. 

Heres the test coverage when we run the command **go test ./... -v -coverpkg=./...**:

 <p align="center">
  <a href="https://git.gvk.idi.ntnu.no/andmag/project_cloud_2019">
    <img src="/images/test_cover.png" alt="Test" width="400" height="50">
  </a>
</p>
 

## Reflection

### What went well
* We worked a lot in the beginning of the project, so we could focus on polishing the final result during the last week. 
* We got to explore the technologies we were curious about when it comes to discord and its support for bots and webhooks. 
* The teamwork worked well, and we all contributed to the project and learned a great deal from each other. 

### What went wrong
* We used a considerably amout of time trying to get files into packages. When we started the project, the only package we had was main, while all other files were "global" in the project module. If we had used a better practice from the beginning, it wouldn't have been so hard structuring things into packages. 
* Related to the one above; We also wanted to have more packages than cmd, src and tests (a database package would've been nice), however since we didn't use it from the beginning, it was hard to adjust our code to that wish. 

## What we have learned
* In this project we learned how to use API's and databases to recieve, process and send data to a bot in a discord channel. 
* We learned how to make use of public libraries for golang such as the discordgo package for the bot and the go-chart package for making the pie charts. We also got slightly better at reading and understanding documentation. 
* We have learned how to utilize docker. 
* We have gotten better at searching for alternative methods for achiving our goals, never given up until we get a result we want. 
* We have gotten more proficient in using go modules/packages for structuring our files. 

## Total work hours
* We logged our hours at school and ended at approximently 120 hours total. We didn't log our hours when we worked individually at home, but we estimate that we ended up with around 20 hours total. Which means the total hours working on the project is around 140 hours total. 
## References
**API's:** 
* [Fox](https://randomfox.ca)
* [Weather](https://www.weatherbit.io/api)

**Libraries:**
* [Discordgo](https://github.com/bwmarrin/discordgo)
* [Embed.go](https://gist.github.com/Necroforger/8b0b70b1a69fa7828b8ad6387ebb3835)
* [Go-chart](https://github.com/wcharczuk/go-chart)

**PaaS:**
* [IFTTT](https://ifttt.com/)