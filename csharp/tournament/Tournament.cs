using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;

public enum Result
{
    win,
    loss,
    draw
}

public class TeamStat
{
    public string Team { get; }

    public int Won { get; private set; }

    public int Lost { get; private set; }

    public int Drawn { get; private set; }

    public int MatchesPlayed => Won + Lost + Drawn;

    public int Points => Won * 3 + Drawn;

    public TeamStat(string team)
    {
        Team = team;
    }

    public void AddWin() => Won++;
    public void AddLoss() => Lost++;
    public void AddDraw() => Drawn++;
}

public static class Tournament
{
    public static void Tally(Stream inStream, Stream outStream)
    {
        string input;
        using (var streamReader = new StreamReader(inStream))
        {
            input = streamReader.ReadToEnd();
        }

        var matchResults = input.Split('\n', StringSplitOptions.RemoveEmptyEntries);
        var statsAsStrings = BuildTeamStats(matchResults);

        using (var streamWriter = new StreamWriter(outStream))
        {
            streamWriter.Write("Team                           | MP |  W |  D |  L |  P");
            statsAsStrings.ForEach(x => streamWriter.Write($"\n{x.Team,-30} | {x.MatchesPlayed,2} | {x.Won,2} | {x.Drawn,2} | {x.Lost,2} | {x.Points,2}"));
        }
    }

    private static List<TeamStat> BuildTeamStats(string[] matchResults)
    {
        var teamStats = new Dictionary<string, TeamStat>();
        
        foreach (var matchResult in matchResults)
        {
            var values = matchResult.Split(';', StringSplitOptions.RemoveEmptyEntries);

            var team1 = values[0];
            var team2 = values[1];
            var result = Enum.Parse(typeof(Result), values[2]);

            if (!teamStats.ContainsKey(team1))
            {
                teamStats.Add(team1, new TeamStat(team1));
            }

            if (!teamStats.ContainsKey(team2))
            {
                teamStats.Add(team2, new TeamStat(team2));
            }

            switch (result)
            {
                case Result.win:
                    teamStats[team1].AddWin();
                    teamStats[team2].AddLoss();
                    break;
                case Result.loss:
                    teamStats[team1].AddLoss();
                    teamStats[team2].AddWin();
                    break;
                case Result.draw:
                    teamStats[team1].AddDraw();
                    teamStats[team2].AddDraw();
                    break;
            }
        }

        return teamStats.Values.OrderByDescending(x => x.Points).ThenBy(x => x.Team).ToList();
    }
}
