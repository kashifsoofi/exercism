using System;

public enum AccountStatus
{
    Opened,
    Closed
}

public class BankAccount
{
    private readonly object balanceLock = new object();

    private decimal balance;
    public AccountStatus status;

    public void Open()
    {
        status = AccountStatus.Opened;
        balance = 0m;
    }

    public void Close()
    {
        status = AccountStatus.Closed;
    }

    public decimal Balance
    {
        get
        {
            if (status == AccountStatus.Closed)
            {
                throw new InvalidOperationException();
            }

            lock (balanceLock)
            {
                return balance;
            }
        }
    }

    public void UpdateBalance(decimal change)
    {
        if (status == AccountStatus.Closed)
        {
            throw new InvalidOperationException();
        }

        lock (balanceLock)
        {
            balance += change;
        }
    }
}
