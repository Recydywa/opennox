#ifndef NOX_PORT_CLIENT_DRAWABLE_DRAWDB
#define NOX_PORT_CLIENT_DRAWABLE_DRAWDB

#include "defs.h"

typedef struct nox_parse_thing_funcs_t {
	const char* name;
	bool (*parse_func)(nox_thing*, nox_memfile*, char*);
} nox_parse_thing_funcs_t;

extern nox_parse_thing_funcs_t nox_parse_thing_funcs[];
extern int nox_parse_thing_funcs_cnt;

void* nox_xxx_parseThingBinClient_44C840_read_things(void);
int nox_parse_thing(nox_memfile* thing_file, char* scratch_buffer, nox_thing* thing);

int nox_xxx_freeObjectTypes_4E2A20();

void sub_4E29D0();

#endif // NOX_PORT_CLIENT_DRAWABLE_DRAWDB
