
# Basic Server and Load Tester

Based on John Crickett's Coding Challenges, I've decided to try one of them out - well, actually two of them.

## Part 2: The Load Tester

This is basically a continuation of my web server challenge project. 

### [Web Server in Go](https://github.com/JamesCalingo/go_server)

We now have a web server that can parse requests from clients and serve them various pieces of data. However, it's important that we're not only able to handle any traffic on our server, but handle it *well*; for example, we wouldn't want to, say, run into massive issues should our server be called upon by [EVERY SINGLE GAMESTOP IN THE UNITED STATES](https://www.polygon.com/2015/4/2/8337499/gamestops-website-down-amiibo-ness). Therefore, something like a load tester/bearer would be good for us to have.

The first step is to make sure we can send/capture any requests that end up on our server. Thankfully, the ol' `http` package helps us immensely here, as we can set up a GET request (this also helps us test out our server as well, since we can now see the 200 OK response from, for example, the first part of the server task). One thing I ran into when doing this step: `http.Get` takes in a url, but said url (at least in my case) REQUIRES the http:// "prefix", so to speak - it panics if this is omitted (I got an "unsupported protocal scheme" error). I should also note that we can make requests to things on the web this way as well; I tried this with popular websites, and most of them worked - with the notable exception of ex-Twitter (see what I did there) which gave me a 400 Bad Request error.

Now that we have a way to send requests, let's send a bunch of them. We're going to start by using a loop to send our requests, as it's the most natural starting place for this IMO. Through this method, I ended up finding out that for my server, it was able to take just over 10,000 requests before it (the server) basically gave up and started spitting out errors ("dial tcp: lookup localhost: no such host"). I thought about trying this online, but that is essentially a DDOS attack: a bunch of "junk" requests to a server, and I'd rather not end up on a list (even if I tried this on my own websites due to how they're hosted), so...

It was at this point where I realized something: I should not be hard coding certain things.

Thanks to how Go handles inputs, you COULD do a "form" wherein you ask for a URL, HTTP method, and a number of times you want to send the requests, but I've found that this can lead to some wild issues.

So now that we can make a (bleep) ton of requests, there's one thing we need to think about: that loop makes those requests sequentially, and more often than not, we're not going to see reqests come in in sequence - they're all going to be coming in concurrently. Again, we don't want to end up in a GameStop situation, so we should make sure we can check massive concurrent loads as well.

Once again, Goroutines are our friend here. 

Now that we've got all of that set up, let's make like baseball and get into some of the advanced stats