#
# generate test data
#

# http://stackoverflow.com/questions/1803628/raw-list-of-person-names
# references census data url's for names

all: dist.all.last dist.femail.first dist.mail.first
	
dist.all.last:
	curl -sL http://www2.census.gov/topics/genealogy/1990surnames/dist.all.last -o $@

dist.femail.first:
	curl -sL http://www2.census.gov/topics/genealogy/1990surnames/dist.female.first -o $@

dist.mail.first:
	curl -sL http://www2.census.gov/topics/genealogy/1990surnames/dist.male.first -o $@
