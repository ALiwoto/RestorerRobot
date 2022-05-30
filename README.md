# <h1 align="middle"> RestorerRobot </h1>

A bot written in golang using gotd library for backing up databases and uploading them to telegram logging channels.

<hr/>

## how to configure the project
Configuration file is called `config.ini`, it has a structure such as this:
```ini
[main]
app_id = 12345
app_hash = abcd
bot_token = 1234324:abcd
field1 = value1
field2 = value2

[SaitamaRobot]
db_url = postgresql://Username:Password@localhost:5432/DatabaseName
backup_interval = 10
some_field = 1233424

[KigyoRobot]
db_url = postgresql://Username:Password@localhost:5432/DatabaseName
backup_interval = 10
some_field = 78797878

[PsychoPass]
db_url = postgresql://Username:Password@localhost:5432/DatabaseName
backup_interval = 15
some_field = 86586786

[AllMightRobot]
db_url = postgresql://Username:Password@localhost:5432/DatabaseName
# backup interval in days
backup_interval = 20
some_field = 86586786

```

Config file contains `main` sections and other sections, every field in the main section
is applied to the whole project (they are globally used), but every field in each own section, is only applied to that project.

You can add as many as project you would like to the config file, there is no limits in that.
The field `db_url` is shouldn't necessary be hosted on localhost, as long as it's valid and accessible, it will be okay.

**WARNING**: all parameters inside of db_url **SHOULD** be url-encoded, otherwise backing up operation will result in failure.

<hr/>

# Scheduled backups

There will be a time interval between backing up each project. For example you can set backup time-interval of `AllMightRobot` project to 20 (days), that way if 20 days passes from last time `AllMightRobot`'s db got backed up, bot will try to take next backup.

This part is still incomplete.

<hr/>

# Force backup

Owners can forcefully make the bot to backup and upload the compressed file to telegram. Please do notice that the file will be also sent to global log channels, if you don't want this, consider using `--private` flag in your command.

Command for forceful backup is as following:

`/forcebackup SaitamaRobot`

or you can also pass it direct db url:

`/forcebackup postgresql://Username:Password@localhost:5432/DatabaseName`

<hr/>

