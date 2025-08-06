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
    const weekOfYear = getWeekOfYear(now);
    let properWeek = weekOfYear - 20;
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
  
  const entry = data.find(item => 
    item.week.toLowerCase() === weekOfSeason.toLowerCase() && 
    item.day.toLowerCase() === dayOfWeek.toLowerCase()
  );
  
  return entry;
}

exports.handler = async (event, context) => {
  console.log('New readings function v1');
  
  try {
    const headers = {
      'Access-Control-Allow-Origin': '*',
      'Access-Control-Allow-Headers': 'Content-Type',
      'Access-Control-Allow-Methods': 'GET, OPTIONS',
      'Content-Type': 'application/json'
    };
    
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
    
    const { tableName, season, week, dayOfWeek } = getTodaysLiturgicalDate();
    
    let data;
    if (tableName === 'year-one') {
      data = require('./dol-year-1.js');
    } else if (tableName === 'year-two') {
      data = require('./dol-year-2.js');
    } else {
      data = require('./dol-year-1.js');
    }
    
    const todaysReadings = findTodaysReadings(data, season, week, dayOfWeek);
    
    if (!todaysReadings) {
      return {
        statusCode: 404,
        headers,
        body: JSON.stringify({ 
          error: 'Today\'s readings not found',
          debug: { tableName, season, week, dayOfWeek }
        })
      };
    }
    
    return {
      statusCode: 200,
      headers,
      body: JSON.stringify(todaysReadings.lessons)
    };
    
  } catch (error) {
    console.error('Function error:', error);
    return {
      statusCode: 500,
      headers: {
        'Access-Control-Allow-Origin': '*',
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({ 
        error: 'Internal server error',
        message: error.message,
        stack: error.stack
      })
    };
  }
};