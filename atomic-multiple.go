package dummy

func f() {
	// START OMIT
	// we may get multiple reports of an entry being down, so we guard
	// against doing things twice.
	if existing, isAlive := alive[entry.addr]; isAlive {
		// we got a report of a down entry after we already successfully
		// reconnected to it. don't throw away a good client.
		if existing != entry {
			continue
		}
	}
	// END OMIT
}
