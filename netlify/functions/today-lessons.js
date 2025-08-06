const fs = require('fs');
const path = require('path');

// Simple liturgical calendar calculation
function getTodaysLiturgicalDate() {
  const now = new Date();
  const dayOfWeek = now.toLocaleDateString('en-US', { weekday: 'long' });
  
  // Episcopal Church starts Year One in even calendar years, Year Two in odd years
  const tableName = now.getFullYear() % 2 === 0 ? 'year-one' : 'year-two';
  
  const month = now.getMonth() + 1; // JavaScript months are 0-indexed
  const day = now.getDate();
  
  let season, week;
  
  // Basic season calculation based on calendar dates
  // This is simplified - actual liturgical calendar is more complex
  if (month === 12 && day >= 25) {
    season = 'christmas';
    week = '1';
  } else if (month === 1 && day <= 6) {
    season = 'christmas';
    week = '1';
  } else if ((month === 1 && day > 6) || month === 2) {
    season = 'epiphany';
    week = '1';
  } else if (month >= 3 && month <= 4) {
    season = 'lent';
    week = '1';
  } else if (month === 5) {
    season = 'easter';
    week = '1';
  } else if (month >= 6 && month <= 11) {
    season = 'after-pentecost';
    // Simplified calculation for after-pentecost
    const weekOfYear = getWeekOfYear(now);
    let properWeek = weekOfYear - 20; // Rough approximation
    if (properWeek < 1) properWeek = 1;
    else if (properWeek > 29) properWeek = 29;
    week = properWeek.toString();
  } else {
    season = 'advent';
    week = '1';
  }
  
  return { tableName, season, week, dayOfWeek };
}

function getWeekOfYear(date) {
  const start = new Date(date.getFullYear(), 0, 1);
  const diff = date - start;
  const oneWeek = 1000 * 60 * 60 * 24 * 7;
  return Math.floor(diff / oneWeek) + 1;
}

function findTodaysReadings(data, season, week, dayOfWeek) {
  let weekOfSeason = `Week of ${week} ${season}`;
  if (season === 'after-pentecost') {
    weekOfSeason = `Proper ${week}`;
  }
  
  // Case-insensitive search
  const entry = data.find(item => 
    item.week.toLowerCase() === weekOfSeason.toLowerCase() && 
    item.day.toLowerCase() === dayOfWeek.toLowerCase()
  );
  
  return entry;
}

exports.handler = async (event, context) => {
  try {
    // Set CORS headers
    const headers = {
      'Access-Control-Allow-Origin': '*',
      'Access-Control-Allow-Headers': 'Content-Type',
      'Access-Control-Allow-Methods': 'GET, OPTIONS',
      'Content-Type': 'application/json'
    };
    
    // Handle preflight requests
    if (event.httpMethod === 'OPTIONS') {
      return { statusCode: 200, headers, body: '' };
    }
    
    if (event.httpMethod !== 'GET') {
      return {
        statusCode: 405,
        headers,
        body: JSON.stringify({ error: 'Method not allowed' })
      };
    }
    
    // Get today's liturgical date
    const { tableName, season, week, dayOfWeek } = getTodaysLiturgicalDate();
    
    // Determine which JSON file to read
    let filename;
    if (tableName === 'year-one') {
      filename = 'dol-year-1.min.json';
    } else if (tableName === 'year-two') {
      filename = 'dol-year-2.min.json';
    } else {
      filename = 'dol-year-1.min.json'; // fallback
    }
    
    // Read the JSON file from the functions directory
    const filePath = path.join(__dirname, filename);
    
    if (!fs.existsSync(filePath)) {
      return {
        statusCode: 500,
        headers,
        body: JSON.stringify({ 
          error: 'Reading data file not found',
          debug: { filePath, tableName, season, week, dayOfWeek, __dirname }
        })
      };
    }
    
    const rawData = fs.readFileSync(filePath, 'utf8');
    const data = JSON.parse(rawData);
    
    // Find today's readings
    const todaysReadings = findTodaysReadings(data, season, week, dayOfWeek);
    
    if (!todaysReadings) {
      return {
        statusCode: 404,
        headers,
        body: JSON.stringify({ 
          error: 'Today\'s readings not found',
          debug: { tableName, season, week, dayOfWeek, weekOfSeason: season === 'after-pentecost' ? `Proper ${week}` : `Week of ${week} ${season}` }
        })
      };
    }
    
    // Return only lessons
    return {
      statusCode: 200,
      headers,
      body: JSON.stringify(todaysReadings.lessons)
    };
    
  } catch (error) {
    console.error('Error:', error);
    return {
      statusCode: 500,
      headers: {
        'Access-Control-Allow-Origin': '*',
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({ 
        error: 'Internal server error',
        message: error.message 
      })
    };
  }
};