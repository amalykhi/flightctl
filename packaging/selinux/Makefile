TARGET?=flightctl_agent
MODULES?=${TARGET:=.pp.bz2}
SHAREDIR?=/usr/share

all: ${MODULES}

%.pp.bz2: %.pp
	rm -f $@ || true
	@echo Compressing $^ -\> $@
	bzip2 -9 $^

%.pp: %.te
	make -f ${SHAREDIR}/selinux/devel/Makefile $@

clean:
	rm -f *~  *.tc *.pp *.pp.bz2
	rm -rf tmp *.tar.gz

man: install-policy
	sepolicy manpage --path . --domain ${TARGET}_t

install-policy: all
	sudo semodule -i ${TARGET}.pp.bz2

install: man
	sudo install -D -m 644 ${TARGET}.pp.bz2    ${DESTDIR}${SHAREDIR}/selinux/packages/
	sudo install -D -m 644 ${TARGET}_selinux.8 ${DESTDIR}${SHAREDIR}/man/man8/
