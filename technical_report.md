# JE-WOTD (Japanese English Word of the Day)

Elijah Hursey

Evangeline Luciano

4/18/2018

# Abstract

The goal of this project is to provide a resource for Japanese people who want to learn English. The purpose of this resource is to provide a new word every day in their native language (Japanese) as well as a word in English for them to study and memorize. In addition to the Japanese and English word, the reading of the Japanese word (in the event that it is kanji) is provided as well as the part(s) of speech for the word in English. This is achieved by using a combination of 2 external resources, namely Jisho and Yandex, and 3 different apps which make up the provided resource. The 3 apps written for this project are the JE-WOTD-api, the email_wkr, and the JE-WOTD-spa (single page application). The results of the combination of these apps and external resources provides a website where a user can view a new word each day, and also an email service which sends the same data to a user’s email should they choose to register.

# Keywords
* Go
* VueJS
* Javascript
* ESLint
* Godeps
* NPM
* Github
* Heroku
* AWS
* DynamoDB
* Jisho
* Yandex
* Japanese
* English
* translation
* api
* spa
* wkr (worker)
* microservice

# Table of Contents

1. [Introduction and Project Overview](#introduction-and-project-overview)
2. [Design, Development, and Testing](#design-development-and-testing)
3. [Results](#results)
4. [Conclusion and Future Work](#conclusions-and-future-work)
5. [References](#references)

# Introduction and Project Overview

I began studying Japanese in May of 2015 and not long after began a 2 semester long study abroad in Japan where I took various language and culture classes. In my time in Japan I made many friends who were studying English just as fervently as I was studying Japanese and after a while I stumbled on a word of the day service that helped me learn many new vocabulary words. Having had such a pleasant experience I began looking for equivalent word of the day services but for my Japanese friends to use and quickly discovered that such a service did not really exist. Thus after a few years I decided I would make my own with objective being to create a user friendly and accurate word of the day service that could help native Japanese speakers begin expanding their English vocabulary. 

All of my Japanese friends used books that were organized by levels with level one being very simple words that you would use very frequently in conversation, such as mother and father. As you buy books of higher levels you get more difficult and less commonly used words. Since these books are so prevalent not many people even consider a solution like a word of the day service. There are obvious pros and cons to using a word of the day solution as opposed to these books. The pros of using a word of the day service is it is very easy to memorize one word every day and can even become a sort of challenge to use it in a conversation that day which can make memorizing the word fun and interactive. The cons would be that you only get a single word every day and if you are trying to learn a lot quickly this would be much too slow. However, used in tandem with a book or class material it provides something unexpected in what can become the mononity of memorizing lots of similar level vocabulary. 

While JE-WOTD provides a word in Japanese, the reading for that word, the English translation, and parts of speech, there are other things that could be added that would give more information. Some of these things would be example sentences in both languages showing both a use case for the Japanese word and the English word. This would provide a whole world of context that can be quintessential in understanding the actual meaning of a word. Another useful feature would be including the dictionary pronunciation of the English word. This could be useful for particularly strange words (e.g. xylophone, Wednesday, February, island), or for users who are not familiar with pronunciation rules in English. 

The full list of features is random word selection, word translation, part of speech retrieval, web page to display data, email registration on website, daily email with content from web page, and dynamic formatting if a word is katakana. The api handles word selection, translation, part of speech retrieval, and email registration while the spa displays data to the user via a web page and the email worker sends out the same content displayed on the web page retrieved from the api every 24 hours.

# Design, Development, and Testing

The design for this project is fairly simple and aims to follow microservice design principles. As such there are the previously stated three major components being the spa, api, and worker. Each of these services is deployed using Heroku. Originally I considered using AWS (Amazon Web Service) to deploy each service, but as I was already familiar with AWS from other projects I decided to use Heroku in an attempt to see what other options there are and how they compare. The api and worker are written in Go and the spa is written in the Javascript framework VueJS. I chose to use these languages and frameworks because I had previous experience with them and wanted to spend the semester making progress on my project instead of trying to stumble along while learning a new language. I also decided that I would be working with someone who had no experience in either of these languages which only further pushed me towards something I had experience with. Godeps was used as the dependency management package for all Go services and npm was used for the spa. ESLint was used to maintain clean javascript in the spa. The most complex of these is the api as it handles data compilation as well as email registration. 

When the api is started the first thing it does check to make sure the emails table is available in DynamoDB. Next it pulls down a list of commonly used English words then it sets a seed based off of the current time in nanoseconds. After the seed is set it randomly gets a word from the list and holds the current time. From this point it simply waits for any incoming http requests on one of it’s four routes. The first route is the basic get word route which checks the current timestamp and if it has not been 24 hours since the last new word it simply returns the one it’s holding. If it has been more than 24 hours it will get a new word and return it. The process of translating a word is that it goes from English, is translated by Yandex, then the translated word is sent to Jisho which comes back with several possible translations and the reading and parts of speech. The second route is a force version of the first route that ignores the time check for debugging purposes. The third handles email registration and simply inserts the provided email into the DynamoDB table. Finally the fourth route scans the DynamoDB table and returns all stored emails. This endpoint is only used by the email worker when it determines who to email when it is triggered.

The email worker much simpler, it is designed such that it makes two requests to the api, one for the list of saved emails and the other for the current word. After these two requests complete it inserts the word data into an html template and sends it to each email address in the list it received from the api. This worker is designed to close when it is done executing as opposed to the api which runs continuously and is triggered by a timer add on in Heroku.

The spa is also quite simple, it simply fills a table with the data received by making a request to the api that is triggered whenever the page is loaded. If the provided data does not have data in the ‘Japanese’ field then it assumes that it is katakana and rearranges the table accordingly. For development purposes there is a new word button that uses the force route in the api to ignore the time restriction. There is also a text box at the top of a page that accepts an email from the user and sends the provided email to the api to be stored in DynamoDB.

The first system that was developed was the api. In the beginning it simply got commonly used words and randomly selected one to be translated via Jisho. In this state each request would return a new word. Once this was working a simple web page was developed that displayed the information provided by the api in a simple table with essentially no styling.

After both of these simple systems were functioning some simple styling was added and the Heroku deployment process was determined. This was much more complicated than expected as for the purpose of this class I was required to keep each individual service in a single Github repository. The final solution to this issue was to create a Github project in each services’ folder that only held a remote attached to Heroku with the main Git project maintaining the actual Github remote. With this setup you have to manually add the Heroku remote to each service if you switch computers which can become inconvenient quickly, but with future work it can be separated out into different repositories.

With deployment was ironed out the Yandex translation middle step was added in to strive for more consistent accurate translations. The thought process behind this was it does not matter what word we started with but if we can get a Japanese word from Yandex and then get the English for that word it is more likely to be accurate than a single translation from only Jisho. Yandex was also chosen because they have a free tier package as opposed to Google Translate which does not.

Now that we were getting consistent translations we added a timing system that would only provide a new word if it had been at least 24 hours since the last new word was retrieved. Previous to this change every time you refreshed the webpage or sent a request to the api you would get a new word and translation. As this change made debugging difficult we also added the force new word route and button in the spa. After that the only thing left was to get the email service working so we wrote the worker that takes a list of emails from DynamoDB and sends an email to each email in the table. 

As the spa and email worker are quite simple testing seemed unnecessary so I focused testing into the api. There are several different data transformations that occur within the api specifically when a new word is selected and translated. The main transformations are when the response from Jisho is subset into a condensed version, actually getting a word to start with, and changing the parts of speech received from Jisho to their Japanese counterparts. 

# Results

The end product ended up being almost exactly what I had in mind when I started thinking about this project. The goal was to have an api that could serve a new word and translation with the parts of speech and reading of the Japanese if it was kanji. The only piece that I could not find a good way to do was having example sentences due to not being able to find a good resource. I also wanted to display these translations on a web page that was hosted in a way that my Japanese friends could use the service with minimal effort. This was achieved with the spa and deployment via Heroku. The last thing I wanted was an email service that would send emails daily so that you would be notified of the new word instead of having to go and check the website every day which was accomplished with the email worker and the DynamoDB table. From a user’s perspective you can navigate to the web page and daily you will see a new word and translation and should you choose you can easily register your email to receive a daily email. Since the early stages of this project I have had one of my friends in Japan using the website and providing me feedback and he has now registered with the email service as well and is actually using it to expand his English vocabulary. This sort of use is exactly what I intended for when I set out to start working on this project.

There were actually very few problems that we encountered working on this project. At the very start there was some delay in starting just because finding a way to do translation that was simple took quite a bit of research. The largest problem was trying to get the different services deployed with Heroku while they all lived in a single Github repository. Another minor issue is that when one of the services has no traffic for a period of time longer than 30 minutes, since we are using the free tier of Heroku, it is shut down which can cause issues with the 24 hour timer set up between words. This is however negligible because I picture more people using the email service than the actual web page and I could upgrade to a version of Heroku that runs continuously very easily should the need arise.

# Conclusions and Future Work

All in all I think that this project was very successful. Aside from having example sentences we were able to accomplish everything we set out to do. There is now a working word of the day service that anyone can use should they want and currently there is at least one native Japanese speaker actively using this project. 

Thinking back on this project I am glad that we decided to use not only the languages that we used but also the other technologies (Heroku, DynamoDB). Go and VueJS are both designed so that you can very quickly get a working product that you can iterate on to improve. This made starting out very easy as we were able to get a working api and spa all in about two weeks time. Initially I did not want to use any AWS services, which was the primary reason I ended up using Heroku, but because of this and the add-ons available through Heroku the email worker was much easier to set up than I originally thought it would be.
Future work that could be done are things like separating each service into its own Github repository, adding support for example sentences, switching to a more reliable translation service, and implementing a mobile friendly view for both the website and emails. The most useful of these would be the mobile friendly version as both the emails and website are practically unreadable without switching your orientation to landscape.

# References

* Go
* Godeps
* VueJS
* NPM
* Github
* Heroku
* AWS
* DynamoDB
* Jisho
* Yandex
