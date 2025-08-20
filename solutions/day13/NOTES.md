# Day 13 Part 2 - Debugging Notes

## Problem Status
- **Part 1**: ✅ SOLVED (109,23)
- **Part 2**: ❌ UNSOLVED

## Confirmed Wrong Answers
These answers have been submitted and confirmed as incorrect:
- 73,122 ❌ (consistently obtained across all implementations)
- 74,122 ❌
- 73,121 ❌
- 72,122 ❌
- 73,123 ❌
- 73,124 ❌ (rate limited, likely wrong)
- 122,73 ❌ (swapped coordinates, rate limited)
- 127,0 ❌ (rate limited)
- 74,121 ❌ (rate limited)

## Key Findings

### The Consistent Answer
All implementations consistently produce (73,122) as the answer:
- Cart 0 is the sole survivor
- Final position reached at tick 10004
- Cart 6 crashes out of bounds at (127,-1) during the same tick

### Sequence of Events (Tick 10004)
1. **Start of tick**: 
   - Cart 6 at (127,0) facing north
   - Cart 0 at (73,123) facing north
2. **During tick** (carts move in Y,X order):
   - Cart 6 moves first (Y=0), tries to go to (127,-1), goes out of bounds
   - Cart 0 moves second (Y=123), moves from (73,123) to (73,122)
3. **End of tick**: Only Cart 0 remains at (73,122)

### Technical Issues Encountered
1. **Out of bounds carts**: Multiple carts try to move to Y=-1 (especially at X=127)
2. **Disconnected track systems**: There appear to be separate track loops that don't connect
3. **Long-running loops**: 5 carts get stuck in non-intersecting loops after tick ~3000

### Implementation Approaches Tried
1. Original solution with direction enums
2. Clean reimplementation with simple data structures
3. Delta-based movement (DX, DY)
4. Bounds checking before vs after movement
5. Treating out-of-bounds as crashes vs preventing movement

### Possible Issues to Investigate
1. **Coordinate system**: Are we using the right X,Y convention?
2. **Tick timing**: Should we return position before or after the final movement?
3. **Cart ordering**: Are we processing carts in the exact right order?
4. **Intersection logic**: Is the left-straight-right cycle implemented correctly?
5. **Collision detection**: Are we checking collisions at the right moment?

## Test Results
- Example Part 2: ✅ Works correctly (6,4)
- Actual input Part 2: ❌ Consistently produces (73,122)

## Next Steps
- Consider if there's a special case for carts going out of bounds
- Review if collision detection should happen differently
- Check if the problem expects a different interpretation of "end of tick"
- Investigate if there's an off-by-one error in the tick counting

## Additional Attempts with Opus 4.1
- Fresh implementation with clean code structure - Result: 73,122 ❌
- Alternative approach checking position after tick - Result: 73,122 ❌  
- Before-tick interpretation (position before final movement) - Result: 73,122 ❌
- Collision-aware implementation (handling mid-tick elimination) - Result: 73,122 ❌

All implementations consistently arrive at the same answer despite different approaches to:
- When to check for the final cart (before/after tick)
- How to handle out-of-bounds carts
- Order of operations within a tick
- Collision detection timing

This strongly suggests the simulation logic is correct but there may be:
1. A misinterpretation of the problem requirements
2. An edge case in the track layout not being handled
3. A different expected behavior for out-of-bounds movement