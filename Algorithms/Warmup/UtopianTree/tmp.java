            if (start.getMicros() <= end.getMicros()) {
                Map<Timestamp, Set<Object>> chronologie = chronologize(key, record);
                Timestamp exactStartTimestamp = Iterables.getFirst(
                        chronologie.keySet(), null);
                Timestamp exactEndTimestamp = Iterables.getLast(chronologie.keySet());
                
                if (exactStartTimestamp != null && 
                        start.getMicros() == end.getMicros()) {
                    if (start.getMicros() >= exactStartTimestamp.getMicros()) {
                        if (start.getMicros() >= exactEndTimestamp.getMicros()) {
                            Entry<Timestamp, Set<Object>> lastEntry = 
                                    Iterables.getLast(chronologie.entrySet());
                            result.put(lastEntry.getKey(), lastEntry.getValue());
                        } else {
                            // if ==, set exact and start adding 1
                            // if >, start binary search and start adding 1
                            // DONE
                            int i = findIndexOfClosestStartTimestamp(chronologie,
                                    start, 0, chronologie.size()-1);
                            Entry<Timestamp, Set<Object>> entry =
                                    Iterables.get(chronologie.entrySet(), i);
                            result.put(entry.getKey(), entry.getValue());     
                        }
                    }
                }
                
                if (exactStartTimestamp != null && 
                        start.getMicros() < end.getMicros()) {
                    if (end.getMicros() > exactStartTimestamp.getMicros()) {
                        Iterator<Entry<Timestamp, Set<Object>>> chronologieIter =
                                chronologie.entrySet().iterator();
                        
                        if (start.getMicros() <= exactStartTimestamp.getMicros()) {
                            // set to first and iterate through loop till end or finish loop
                            // DONE
                            while(chronologieIter.hasNext()) {
                                Entry<Timestamp, Set<Object>> entry = 
                                        chronologieIter.next();
                                if (entry.getKey().getMicros() < end.getMicros()) {
                                    result.put(entry.getKey(), entry.getValue());
                                } else {
                                    break;
                                }
                            }
                            
                        } else {
                            
                            if (start.getMicros() >= exactEndTimestamp.getMicros()) {
                                // binary search to find closest/ exact 
                                // and add iteration to end or finish loop
                            } else {
                                
                            }

                        }
                    }
                }
            }
            return result;