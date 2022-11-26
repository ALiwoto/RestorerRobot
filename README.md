# <h1 align="middle"> RestorerRobot </h1>

A bot written in golang using mtproto for backing up databases and uploading them to telegram logging channels.
Telegram's max limit for uploading files is 2GB, which means the compressed file (.zip)'s size that bot is trying to upload to telegram should not be more than 2GB.

<hr/>

## how to configure the project
Configuration file is called `config.ini`, it has a structure such as this (for a more completed sample file, please visit [config.sample.ini](config.sample.ini) file):
```ini
[main]
app_id = 12345
app_hash = abcd
bot_token = 1234324:abcd
backups_base_directory_path = backups

[SaitamaRobot]
db_url = postgresql://Username:Password@localhost:5432/DatabaseName
backup_interval = 10
# additional log channels, backup file will be sent to these channels
log_channels = -10012548, -1005487
backup_type = dump

[KigyoRobot]
db_url = postgresql://Username:Password@localhost:5432/DatabaseName
backup_interval = 10
# additional log channels, backup file will be sent to these channels
log_channels = -10012548, -1005487
backup_type = dump

[PsychoPass]
db_url = postgresql://Username:Password@localhost:5432/DatabaseName
backup_interval = 15
# additional log channels, backup file will be sent to these channels
log_channels = -10012548, -1005487
backup_type = dump

[AllMightRobot]
db_url = postgresql://Username:Password@localhost:5432/DatabaseName
# backup interval in days
backup_interval = 20
# additional log channels, backup file will be sent to these channels
log_channels = -10012548, -1005487
backup_type = dump

[StalkerGameRobot]
db_path = E:\StalkerGameRobot\user.db
# backup interval in days
backup_interval = 7
#log_channels = -10012548, -1005487
backup_type = sqlite
```

Config file contains `main` sections and other sections, every field in the main section
is applied to the whole project (they are globally used), but every field in each own section, is only applied to that project.

You can add as many as project you would like to the config file, there is no limits in that.
The field `db_url` is shouldn't necessary be hosted on localhost, as long as it's valid and accessible, it will be okay.

**WARNING**: all parameters inside of db_url **SHOULD** be url-encoded, otherwise backing up operation will result in failure.

<hr/>

### Log channels
There are two types of log channels: global log channels and separated log channels.
All backup files will be sent to global log channels, regardless of when, who and why the bot is taking backup. Those who have access to those log channels will be able to download all of backup files shared by this bot.

Separated log channels are related to their own project, another projects' backup files will not be shared in those log channels. For example consider bot is trying to get backup from SaitamaRobot's database, the compressed file will be sent to log-channel and SaitamaRobot's separated log-channel (`log_channels` config variable can point to a channel, group or a user).

In the case a user uses `/forcebackup` command, the compressed file will be sent to log-channels AND the user's PM.

<hr/>

### Scheduled backups

There will be a time interval between backing up each project. For example you can set backup time-interval of `AllMightRobot` project to 20 (days), that way if 20 days passes from last time `AllMightRobot`'s db got backed up, bot will try to take next backup automatically.
And then the compressed backup file will be sent to the log channel.

This part is still incomplete.

<hr/>

### Backup types
This bot currently supports 2 types of backup file output:
- .dump
- .sql

You can set `backup_type` variable in `config.ini` file to either of them (please don't include '.').
Default backup type is set to `sql`.

<hr/>

### Force backup

Owners can forcefully make the bot to backup and upload the compressed file to telegram. Please do notice that the file will be also sent to global log channels, if you don't want this, consider using `--private` flag in your command.

Command for forceful backup is as following:

`/forcebackup SaitamaRobot`

or you can also pass it direct db url:

`/forcebackup postgresql://Username:Password@localhost:5432/DatabaseName`

<hr/>


<h2 align="middle"> 
    - By <a href=https://t.me/Kaizoku> Kaizoku </a>
</h2>

