CREATE DATABASE news 
USE news

CREATE TABLE feed (
  provider TEXT NOT NULL,
  category TEXT NOT NULL,
  url TEXT NOT NULL
);

INSERT INTO feed VALUES ('bbc', 'uk', 'http://feeds.bbci.co.uk/news/uk/rss.xml');
INSERT INTO feed VALUES ('bbc', 'tech', 'http://feeds.bbci.co.uk/news/technology/rss.xml');
INSERT INTO feed VALUES ('sky', 'uk', 'http://feeds.skynews.com/feeds/rss/uk.xml');
INSERT INTO feed VALUES ('sky', 'tech', 'http://feeds.skynews.com/feeds/rss/technology.xml');
