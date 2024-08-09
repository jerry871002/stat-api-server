import React from 'react';
import LineupSlot from './LineupSlot';

const Lineup = ({ lineup, movePlayerToSlot, removePlayerFromSlot }) => {
  return (
    <div>
      <h2>Lineup</h2>
      {lineup.map((player, index) => (
        <LineupSlot
          key={index}
          index={index}
          player={player}
          movePlayerToSlot={movePlayerToSlot}
          removePlayerFromSlot={removePlayerFromSlot}
        />
      ))}
    </div>
  );
};

export default Lineup;