import { useState, useEffect } from 'react';
import { Container, TextField, Button, Typography, Box, Paper, CircularProgress, Autocomplete, Grid } from '@mui/material';
import { PictureAsPdf } from '@mui/icons-material';
import axios from 'axios';
import { Toaster, toast } from 'react-hot-toast';
import { jsPDF } from 'jspdf';
import { format } from 'date-fns';

interface ExamSchedule {
  Dept: string;
  CourseCode: string;
  CourseTitle: string;
  Section: string;
  Teacher: string;
  ExamDate: string;
  ExamTime: string;
  Room: string;
}

function App() {
  const [userId, setUserId] = useState('');
  const [courseCount, setCourseCount] = useState('');
  const [courses, setCourses] = useState<Array<{ title: string; section: string }>>([]);
  const [loading, setLoading] = useState(false);
  const [schedules, setSchedules] = useState<ExamSchedule[]>([]);
  const [courseTitles, setCourseTitles] = useState<string[]>([]);

  useEffect(() => {
    const fetchCourseTitles = async () => {
      try {
        const response = await axios.get('http://localhost:7070/exam-schedule');
        const uniqueTitles = Array.from(new Set(response.data.map((exam: ExamSchedule) => exam.CourseTitle))) as string[];
        setCourseTitles(uniqueTitles.sort((a, b) => a.localeCompare(b)));
      } catch (error: any) {
        console.error('Failed to fetch course titles:', error);
        if (error.response) {
          // Server responded with an error
          toast.error(`Failed to load course suggestions: ${error.response.data || 'Server error'}`);
        } else if (error.request) {
          // Request was made but no response
          toast.error('Server is not responding. Please check if the server is running.');
        } else {
          // Error in request setup
          toast.error('Failed to load course suggestions. Please try again later.');
        }
      }
    };
    fetchCourseTitles();
  }, []);

  const handleCourseCountChange = (value: string) => {
    const count = parseInt(value) || 0;
    if (count > 10) {
      toast.error('Maximum 10 courses allowed');
      return;
    }
    setCourseCount(value);
    setCourses(Array(count).fill({ title: '', section: '' }));
  };

  const handleSubmit = async () => {
    // Validate user ID
    if (!userId.trim()) {
      toast.error('Please enter a User ID');
      return;
    }

    if (!/^\d+$/.test(userId)) {
      toast.error('User ID must be numeric');
      return;
    }

    // Validate course count
    const count = parseInt(courseCount);
    if (!courseCount || isNaN(count) || count <= 0) {
      toast.error('Please enter a valid number of courses');
      return;
    }

    // Validate courses
    if (courses.some(c => !c.title.trim() || !c.section.trim())) {
      toast.error('Please fill in all course titles and sections');
      return;
    }

    setLoading(true);
    try {
      const response = await axios.post('http://localhost:7070/exam-schedule', {
        userId: userId.trim(),
        courses: courses.map(c => ({
          title: c.title.trim(),
          section: c.section.trim()
        }))
      });

      if (response.data.length === 0) {
        toast.error('No exam schedules found for the given courses');
        setSchedules([]);
      } else {
        setSchedules(response.data);
        toast.success(`Found ${response.data.length} exam schedule(s)`);
      }
    } catch (error: any) {
      if (error.response) {
        // Server responded with an error
        toast.error(error.response.data || 'Failed to fetch exam schedules');
      } else if (error.request) {
        // Request was made but no response
        toast.error('Server is not responding. Please try again later.');
      } else {
        // Error in request setup
        toast.error('An error occurred while making the request');
      }
      console.error('Error:', error);
    } finally {
      setLoading(false);
    }
  };

  const generatePDF = () => {
    const doc = new jsPDF();
    const pageWidth = doc.internal.pageSize.getWidth();
    
    // Title
    doc.setFontSize(20);
    doc.text('Exam Schedule', pageWidth/2, 20, { align: 'center' });
    
    // Group schedules by date
    const groupedSchedules = schedules.reduce<Record<string, ExamSchedule[]>>((acc, schedule) => {
      const date = schedule.ExamDate;
      if (!acc[date]) acc[date] = [];
      acc[date].push(schedule);
      return acc;
    }, {} as Record<string, ExamSchedule[]>);

    let yPos = 40;
    
    Object.entries(groupedSchedules).forEach(([date, dateSchedules]) => {
      if (yPos > 250) {
        doc.addPage();
        yPos = 20;
      }
      
      doc.setFontSize(14);
      doc.text(date, 10, yPos);
      yPos += 10;
      
      dateSchedules.forEach((schedule: ExamSchedule) => {
        if (yPos > 250) {
          doc.addPage();
          yPos = 20;
        }
        
        doc.setFontSize(12);
        doc.text(`${schedule.CourseCode} - ${schedule.CourseTitle}`, 15, yPos);
        yPos += 7;
        doc.setFontSize(10);
        doc.text(`Time: ${schedule.ExamTime} | Room: ${schedule.Room}`, 20, yPos);
        yPos += 10;
      });
      
      yPos += 5;
    });
    
    doc.save('exam-schedule.pdf');
  };

  return (
    <Container maxWidth="lg" sx={{ 
      py: 4,
      minHeight: '100vh',
      background: 'linear-gradient(135deg, #1a237e 0%, #0d47a1 100%)',
      display: 'flex',
      flexDirection: 'column',
      gap: 4
    }}>
      <Toaster position="top-right" />
      <Typography variant="h3" component="h1" gutterBottom align="center" sx={{ 
        color: '#fff',
        fontWeight: 'bold',
        textShadow: '2px 2px 4px rgba(0,0,0,0.3)'
      }}>
        Exam Schedule Finder
      </Typography>
      
      <Paper elevation={8} sx={{ 
        p: 4,
        mb: 4,
        background: 'rgba(255, 255, 255, 0.9)',
        backdropFilter: 'blur(10px)',
        borderRadius: 3,
        boxShadow: '0 8px 32px rgba(0,0,0,0.1)'
      }}>
        <Box component="form" noValidate autoComplete="off" sx={{ display: 'flex', flexDirection: 'column', gap: 3 }}>
          <TextField
            fullWidth
            label="User ID"
            value={userId}
            onChange={(e) => setUserId(e.target.value)}
            required
            variant="outlined"
            sx={{ '& .MuiOutlinedInput-root': { borderRadius: 2 } }}
          />
          
          <TextField
            fullWidth
            label="Number of Courses"
            type="number"
            value={courseCount}
            onChange={(e) => handleCourseCountChange(e.target.value)}
            required
            variant="outlined"
            sx={{ '& .MuiOutlinedInput-root': { borderRadius: 2 } }}
          />
          
          {courses.map((course, index) => (
            <Box key={index} sx={{ 
              p: 3,
              borderRadius: 2,
              bgcolor: 'rgba(255,255,255,0.7)',
              boxShadow: '0 4px 12px rgba(0,0,0,0.05)'
            }}>
              <Typography variant="h6" gutterBottom sx={{ color: '#1a237e' }}>
                Course {index + 1}
              </Typography>
              <Autocomplete
                fullWidth
                options={courseTitles}
                value={course.title}
                onChange={(_, newValue) => {
                  const newCourses = [...courses];
                  newCourses[index].title = newValue || '';
                  setCourses(newCourses);
                }}
                renderInput={(params) => (
                  <TextField
                    {...params}
                    label="Course Title"
                    required
                    variant="outlined"
                    sx={{ mt: 2, '& .MuiOutlinedInput-root': { borderRadius: 2 } }}
                    error={!course.title.trim() && courses.length > 0}
                    helperText={!course.title.trim() && courses.length > 0 ? 'Course title is required' : ''}
                  />
                )}
              />
              <TextField
                fullWidth
                label="Section"
                value={course.section}
                onChange={(e) => {
                  const newCourses = [...courses];
                  newCourses[index].section = e.target.value;
                  setCourses(newCourses);
                }}
                required
                variant="outlined"
                sx={{ mt: 2, '& .MuiOutlinedInput-root': { borderRadius: 2 } }}
              />
            </Box>
          ))}
          
          <Button
            fullWidth
            variant="contained"
            onClick={handleSubmit}
            disabled={loading}
            sx={{ 
              mt: 3,
              py: 2,
              bgcolor: '#1a237e',
              borderRadius: 2,
              '&:hover': { bgcolor: '#0d47a1' },
              transition: 'all 0.3s ease'
            }}
          >
            {loading ? <CircularProgress size={24} /> : 'Find Schedules'}
          </Button>
        </Box>
      </Paper>

      {schedules.length > 0 && (
        <Paper elevation={8} sx={{ 
          p: 4,
          background: 'rgba(255, 255, 255, 0.9)',
          backdropFilter: 'blur(10px)',
          borderRadius: 3
        }}>
          <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', mb: 3 }}>
            <Typography variant="h5" sx={{ color: '#1a237e', fontWeight: 'bold' }}>
              Found Schedules
            </Typography>
            <Button
              variant="contained"
              startIcon={<PictureAsPdf />}
              onClick={generatePDF}
              sx={{ 
                bgcolor: '#1a237e',
                '&:hover': { bgcolor: '#0d47a1' }
              }}
            >
              Download PDF
            </Button>
          </Box>
          
          {Object.entries(schedules.reduce<Record<string, ExamSchedule[]>>((acc, schedule) => {
            const date = schedule.ExamDate;
            if (!acc[date]) acc[date] = [];
            acc[date].push(schedule);
            return acc;
          }, {} as Record<string, ExamSchedule[]>)).map(([date, dateSchedules]) => (
            <Box key={date} sx={{ mb: 4 }}>
              <Typography variant="h6" sx={{ 
                color: '#1a237e',
                borderBottom: '2px solid #1a237e',
                pb: 1,
                mb: 2
              }}>
                {date}
              </Typography>
              <Box sx={{ display: 'flex', flexDirection: 'column', gap: 2 }}>
                {dateSchedules.map((schedule, index) => (
                  <Box key={index} sx={{ 
                    p: 3,
                    bgcolor: 'rgba(255,255,255,0.7)',
                    borderRadius: 2,
                    boxShadow: '0 4px 12px rgba(0,0,0,0.05)',
                    transition: 'transform 0.2s ease',
                    '&:hover': { transform: 'translateY(-2px)' }
                  }}>
                    <Grid container spacing={2}>
                      <Grid item xs={12} md={4}>
                        <Typography variant="subtitle1" sx={{ fontWeight: 'bold', color: '#1a237e' }}>
                          {schedule.CourseCode}
                        </Typography>
                        <Typography variant="body1">{schedule.CourseTitle}</Typography>
                      </Grid>
                      <Grid item xs={12} md={4}>
                        <Typography variant="body2" sx={{ color: '#666' }}>
                          Section: {schedule.Section}
                        </Typography>
                        <Typography variant="body2" sx={{ color: '#666' }}>
                          Teacher: {schedule.Teacher}
                        </Typography>
                      </Grid>
                      <Grid item xs={12} md={4}>
                        <Typography variant="body2" sx={{ color: '#666' }}>
                          Time: {schedule.ExamTime}
                        </Typography>
                        <Typography variant="body2" sx={{ color: '#666' }}>
                          Room: {schedule.Room}
                        </Typography>
                      </Grid>
                    </Grid>
                  </Box>
                ))}
              </Box>
            </Box>
          ))}
        </Paper>
      )}
    </Container>
  );
}

export default App;
