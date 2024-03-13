import java.util.Scanner;
class List
{
    int numbers[];

    public void input()
    {
		Scanner sc = new Scanner(System.in);
		int range;
    	System.out.println("Enter the number of numbers you want to enter");
    	range=sc.nextInt();
    	this.numbers=new int[range];
    	int i,j=1;
    	for(i=0;i<range;i++)
    	{
    		System.out.print("Enter the "+(j++)+" number: ");
    		numbers[i]=sc.nextInt();
    	}
    }
    public int getSum()
    {
		int sum = 0;
        for(int i=0;i<numbers.length;i++)
        {
        	sum = sum+numbers[i];
        }
        // System.out.println("The required sum is "+sum);
		return sum;
    }
    public double getAvg()
    {
    	double average=this.getSum()/numbers.length;
    	// System.out.println("Average "+average);
		return average;
    }
	public void Search(int num)
	{
		int i = 0;
		boolean f = false;
		for(i=0;i < numbers.length;i++)
		{
			if(num==numbers[i])
			{
				f=true;
				break;
			}
		}
		if(f==true)
		    System.out.println("The "+ i +" index number is "+num);
		else
			System.out.println("The number is not on the list");
	}
	public double getMax()
	{	
		double max=Double.NEGATIVE_INFINITY;
		for(int i=0;i<numbers.length;i++)
		{
			for(int k=i+1;k<numbers.length;k++)
			{
				if(numbers[i]>max)
				{
					max=numbers[i];
				}
				if(numbers[k]>max)
				{
					max=numbers[k];
				}
			}
		}
		return max;
	}
	public double getMin()
	{
		double min=Double.POSITIVE_INFINITY;
		for(int i=0;i<numbers.length;i++)
		{
			for(int k=i+1;k<numbers.length;k++)
			{
				if(numbers[i]<min)
				{
					min=numbers[i];
				}
				if(numbers[k]<min)
				{
					min=numbers[k];
				}
			}
		}
		return min;
	}
	public void sortAscend()
	{
		int position;
		int temporary;
		int i,k,min;
		for(i = 0;i<numbers.length;i++)
		{
			position = i;
			min = numbers[i];
			for(k = i + 1; k < numbers.length; k++)
			{
				if(numbers[k] < min)
				{
					min = numbers[k];
					position = k;
				}
			}
			temporary = numbers[i];
			numbers[i] = numbers[position];
			numbers[position] = temporary;
		}
		System.out.println("The numbers in ascending order----->");
		for(i = 0; i < numbers.length; i++)
		{
			System.out.println(numbers[i]+"\t");
		}
	}
}
public class TestList 
{
    public static void main(String[] args)
    {
        Scanner a = new Scanner(System.in);
    	List sc = new List();
    	
    	sc.input();
    	
    	boolean f = true;
    	while(f != false)
    	{
    		System.out.println("Enter: \n 1-To Search \n 2-To Sum \n 3-To Average \n 4-To Find Maximum \n 5-To Find the Minimum \n 6-To Sort \n 7-To Terminate ");
        	int ch = a.nextInt();
    	 switch(ch)
    	 {
    	  case 1:
		  	System.out.print("Enter the number you want to search: ");
		  	int num = a.nextInt();
    		sc.Search(num);
    		break;
    	  case 2:
    		int sum = sc.getSum();
			System.out.println("The required sum is " + sum);
    		break;
    	  case 3:
			System.out.println("Average "+ sc.getAvg());
    		break;
    	  case 4:
			System.out.println("The maximum number is "+sc.getMax());
    		break;
    	  case 5: 
			System.out.println("The minimum number is " + sc.getMin());
    		break;
    	  case 6:
    		sc.sortAscend();
    		break;
    	  case 7:
    		f = false;
    		break;
    	  default:
    			System.out.println("Wrong choice. ");
    	}
    	}
    	a.close();
    }
}

