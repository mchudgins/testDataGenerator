#
# test data generator
#

# http://stackoverflow.com/questions/1803628/raw-list-of-person-names
# references census data url's for names

all: census_data lastNames

census_data: dist.all.last dist.female.first dist.mail.first

dist.all.last:
	curl -sL http://www2.census.gov/topics/genealogy/1990surnames/dist.all.last -o $@

dist.female.first:
	curl -sL http://www2.census.gov/topics/genealogy/1990surnames/dist.female.first -o $@

dist.male.first:
	curl -sL http://www2.census.gov/topics/genealogy/1990surnames/dist.male.first -o $@

lastNames: dist.all.last
	awk '{ print $$1 }' $^ >$@.orig
	awk '{ len = length( $$1 ) - 2; last = substr( $$1, 1, len); printf "%sZQ\n", last; }' $^ >$@

femaleNames: dist.female.first
	awk '{ print $$1 }' $^ >$@

maleNames: dist.male.first
	awk '{ print $$1 }' $^ >$@

