# tg-demo-bot
Telegram demo bot

## Installing

1. Create .env file with environment variables
    - `BOT_TOKEN` - token for your bot, you can know ask
    - `DB_DSN` - database dsn
2. Run docker-compose configuration   
```bash
docker-compose up -d
```
3. If previous step was successful, you can see in logs and database


## TODO
   - Fill more fields about messages
   - Implement statistic command
   - Use reindexer/reindexer for work with statistics
   - Refactoring code
   - Implement work with telegram hooks

