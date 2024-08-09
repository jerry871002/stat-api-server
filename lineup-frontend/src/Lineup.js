import React, { useState } from 'react';
import LineupSlot from './LineupSlot';

const Lineup = () => {
  const [lineup, setLineup] = useState(Array(9).fill(null));

  const movePlayerToSlot = (player, index) => {
    const newLineup = [...lineup];
    newLineup[index] = player;
    setLineup(newLineup);
  };

  return (
    <div>
      <h2>Lineup</h2>
      {lineup.map((player, index) => (
        <LineupSlot key={index} index={index} player={player} movePlayerToSlot={movePlayerToSlot} />
      ))}
    </div>
  );
};

export default Lineup;