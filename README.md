# JEWOTD

This project aims to provide a Word of the Day service for native Japanese speakers wanting to learn English.

## Feature List

* A:
  * Parse the EDICT dictionary file for the Japanese word, English word, and part of speech
  * Webpage that serves a new word with the Japanese and English every day as well as the part of speach
  * API that handles providing the webpage with the words
* B:
  * Incorporate Bootstrap or some equivalent styling package in the front end so that it looks nice
  * Display example sentences in both languages for the word
* C:
  * 
  * Add an email service that will send the same data represented on the website to a user each day

## Use case

* The primary intended use case for this product is for native Japanese speakers to have access to a new English word every day.

## Similar Experience

* I have worked in VueJS and Go before in a similar vein where I was responsible for creating an API that provided data to a webpage that I was also responsible for designing. I have not however worked with any sort of email or subscription based service.

## Technology

* Go
* VueJS
* JMdict
* Database (TBD)
* Email service (TBD)

## Risk Areas

* How to host both the API and the Webpage so that they are actually available on the internet.
* Which database will best suit my needs for either storing the JMdict data and/or user emails for the email service
* How to set up the email service to get my data and send to all users