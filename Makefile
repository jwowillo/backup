.PHONY: doc

doc:
	@echo making doc
	pandoc doc/requirements.md --latex-engine xelatex \
		-o doc/requirements.pdf
